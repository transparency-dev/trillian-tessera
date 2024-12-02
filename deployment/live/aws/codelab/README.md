# AWS codelab deployment

This codelab helps you bring a test Trillian Tessera stack on AWS,
and to use it running a test personality server on an EC2 VM.
The Tessera test stack will be comprised of an Aurora RDS MySQL database
and a private S3 bubket. This codelab will also guide you to connect both
the RDS instance and the S3 bucket to your VM.
 
## Prerequisites
For the remainder of this codelab, you'll need to have an AWS account,
with a running EC2 Amazon Linux VM, and the following software installed:
 - [golang](https://go.dev/doc/install), which we'll use to compile and
   run the test personality on the VM
 - [terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
   and [terragrunt](https://terragrunt.gruntwork.io/docs/getting-started/install/)
   in order to deploy the Trillian Tessera stack from the VM.
 - `git` to clone the repo
 - a terminal multiplexer of your choice for convenience

## Instructions

 1. SSH to your VM
 1. Authenticate with a role that has sufficient access to create resources.
    For the purpose of this codelab, and for ease of demonstration, we'll use
    the `AdministratorAccess` role, and authenticate with `aws configure sso`.
    DO NOT use this role to run any production software, or if there are other
    services running on your AWS account.
    Here's an example run:
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
 1. Set these environment variables according to the you chose when configuring
    your AWS profile:
    ```
    export AWS_REGION=us-east-1
    export AWS_PROFILE=AdministratorAccess-<REDACTED>
    ```
 1. Fetch the Tessera repo:
    ```
    git clone https://github.com/transparency-dev/trillian-tessera
    ```
 1. Init terragrunt:
    ```
    terragrunt init --terragrunt-working-dir=deployment/live/aws/codelab/
    ```
 1. Apply everything:
    ```
    terragrunt apply --terragrunt-working-dir=deployment/live/aws/codelab/
    ```
    This brings up the Terraform infrastructure (S3 bucket + dynanoDB table
    for locking) and the Trillian Tessera stack: an RDS Aurora instance,
    a private S3 bucket, and connects this bucket to the default VPC.
 1. Save the RDS instance URI and S3 bucket name for later:
    ```
    export LOG_RDS_DB=$(terragrunt output --terragrunt-working-dir=deployment/live/aws/codelab/ --raw log_rds_db)
    export LOG_BUCKET=$(terragrunt output --terragrunt-working-dir=deployment/live/aws/codelab/ --raw log_bucket_id)
    export LOG_NAME=$(terragrunt output --terragrunt-working-dir=deployment/live/aws/codelab/ --raw log_name)
    ```
 1. Configure the VM and RDS instance to be able to speak to one another following
    [these instructions](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/tutorial-ec2-rds-option3.html),
    it takes a few clicks in the UI.
 1. Generate the key pair used to sign and verify checkpoints:
    ```
    mkdir -p /home/ec2-user/tessera-keys
    go run github.com/transparency-dev/serverless-log/cmd/generate_keys@80334bc9dc573e8f6c5b3694efad6358da50abd4 \
              --key_name=$LOG_NAME/test/conformance \
              --out_priv=/home/ec2-user/tessera-keys/$LOG_NAME.sec \
              --out_pub=/home/ec2-user/tessera-keys/$LOG_NAME.pub
    ```
 1. Run the conformance binary in `trillian-tessera/cmd/conformance/aws`.
    This binary is a small personality that accepts `add/` requests,
    and stores the data in the Trillian Tessera infrastructure you've
    just brought up:
    ```
    go run ./cmd/conformance/aws --bucket=$LOG_BUCKET --db_user=root --db_password=password --db_name=tessera --db_host=$LOG_RDS_DB --signer=$(cat /home/ec2-user/tessera-keys/$LOG_NAME.sec)  -v=3
    ```
 1. Congratulations, you've now successfully brought up a Trillian Tessera
    stack on AWS, and started a personality server that can add entries to it.
    This server accepts `add/` requests at `WRITE_URL=http://localhost:2024/`.
    Log entries can be read directly from S3 without going through the server,
    at `READ_URL=https://$LOG_BUCKET.s3.$AWS_REGION.amazonaws.com/`
 1. Head over to the [remainder of this codelab](https://github.com/transparency-dev/trillian-tessera/tree/main/cmd/conformance#codelab)
    add leaves to the log.
 