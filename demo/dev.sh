# Unless required by ap
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

BASEDIR=$(dirname "$0")

kind create cluster --config "$BASEDIR/kind-cluster-with-extramounts.yaml"

