# AWS Test Configs

Work in progress.


 1. SSH to a VM
 2. Install golang
 3. Install git
 4. Install terragrunt?
 5. Install terragrunt
 6. Install Terminal multiplexer
 7. run `aws configure sso`. Here's an example run:
    ```
    [ec2-user@ip-172-31-21-186 trillian-tessera]$ aws configure sso
    SSO session name (Recommended): greenfield-session
    SSO start URL [None]: https://console.aws.amazon.com/ // unless you use a custom signin console
    SSO region [None]: us-east-1
    SSO registration scopes [sso:account:access]:
    Attempting to automatically open the SSO authorization page in your default browser.
    If the browser does not open or you wish to use a different device to authorize this request, open the following URL:
    
    https://device.sso.us-east-1.amazonaws.com/
    
    Then enter the code:
    
    <REDACTED>
    There are 4 AWS accounts available to you.
    Using the account ID <REDACTED>
    The only role available to you is: AdministratorAccess
    Using the role name "AdministratorAccess"
    CLI default client Region [None]: us-east-1
    CLI default output format [None]:
    CLI profile name [AdministratorAccess-<REDACTED>]:
    
    To use this profile, specify the profile name using --profile, as shown:
    
    aws s3 ls --profile AdministratorAccess-<REDACTED>
    ```
 8. Set environment variables:
    `export AWS_REGION=us-east-1`
    `export AWS_PROFILE=AdministratorAccess-<REDACTED>`
 9. `git clone https://github.com/transparency-dev/trillian-tessera`
 10. `cd trillian-tessera/deployment/live/aws/codelab` 
 11. `terragrunt init`
 12. `terragrunt apply`
 13. save variables for later
     ```
     export LOG_BUCKET=$(terragrunt output -raw log_bucket_id)
     export LOG_RDS_DB=$(terragrunt output -raw log_rds_db)
     ```
 14. Use the UI to connect the VM to the DB instance: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/tutorial-ec2-rds-option3.html
 15. generate keys
     ```
     mkdir -p ~/tessera-keys
     go run github.com/transparency-dev/serverless-log/cmd/generate_keys@80334bc9dc573e8f6c5b3694efad6358da50abd4 \
               --key_name=$TESSERA_PREFIX_NAME-$TESSERA_BASE_NAME/test/conformance \
               --out_priv=/home/ec2-user/tessera-keys/key.sec \
               --out_pub=/home/ec2-user/tessera-keys/key.pub
     ```
 16. Run the conformance binary from within `trillian-tessera/cmd/conformance/aws`
     ```
     go run ./ --bucket=$LOG_BUCKET --db_user=root --db_password=password --db_name=tessera --db_host=$LOG_RDS_DB --signer=$(cat /home/ec2-user/tessera-keys/key.sec)  -v=3
     ```
 17. export WRITE_URL=http://localhost:2024/
 18. export READ_URL=https://$LOG_BUCKET.s3.$AWS_REGION.amazonaws.com/
 19. Follow the codelab: https://github.com/transparency-dev/trillian-tessera/tree/main/cmd/conformance#codelab to send leaves.
