// See the License for the specific language governing permissions and
// limitations under the License.

package ctfe

//go:generate sh -c "protoc -I=. -I$(go list -f '{{ .Dir }}' github.com/google/trillian) -I$(go list -f '{{ .Dir }}' github.com/transparency-dev/trillian-tessera/personalities/ct-static-api) --go_out=paths=source_relative:. configpb/config.proto"
