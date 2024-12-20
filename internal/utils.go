package internal

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func Default[T any](object *T, defaultValue T) T {
	if object == nil {
		return defaultValue
	}
	return *object

}

func Ptr[T any](v T) *T {
	return &v
}

func Val[VAL any, PTR *VAL](ptr PTR) VAL {
	if ptr == nil {
		val := new(VAL)
		return *val
	} else {
		return *ptr
	}
}

func Flatten(m map[string]interface{}) map[string][]byte {
	fmt.Print("Test change for PR")
	o := make(map[string][]byte)
	for k, v := range m {
		k = strings.Replace(k, " ", "_", -1)
		switch child := v.(type) {
		case map[string]interface{}:
			nm := Flatten(child)
			for nk, nv := range nm {
				nk = strings.Replace(nk, " ", "_", -1)
				o[k+"."+nk] = nv
			}
		case string:
			o[k] = []byte(child)
		case int:
			o[k] = []byte(strconv.Itoa(child))
		case float64:
			o[k] = []byte(fmt.Sprintf("%f", child))
		case []byte:
			o[k] = child
		default:
			panic("unhandled json type.")
		}
	}
	return o
}

func SliceEqualsFold(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for _, value := range left {
		if !SliceContainsIgnoreCase(right, value) {
			return false
		}
	}
	return true
}

func SliceContainsIgnoreCase(slice []string, v1 string) bool {
	found := false
	for _, v2 := range slice {

		if strings.EqualFold(v1, v2) {
			found = true
		}
	}
	return found
}

func Float32PtrToIntPtr(float *float32) *int {
	if float == nil {
		return nil
	}
	var val = int(*float)
	return &val
}

// ParseConnectionDetailsFromKubeYaml extracts server Url and certificate bundle from kubeconfig and returns it, if parsing fails empty strings are returned
func ParseConnectionDetailsFromKubeYaml(kubeConfig []byte) (string, string) {
	kubeYaml := &kubeConfigYaml{}
	yamlErr := yaml.Unmarshal(kubeConfig, kubeYaml)

	// gracefully ignore if format isn't matching for now
	if yamlErr != nil || len(kubeYaml.Clusters) == 0 {
		return "", ""
	}
	return kubeYaml.Clusters[0].Cluster.Server, kubeYaml.Clusters[0].Cluster.CertificateAuthData
}

// CopyMaps helper for copying map contents
func CopyMaps[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		dst[k] = v
	}
}

type kubeConfigYaml struct {
	Clusters []namedClusterYaml `json:"clusters"`
}
type namedClusterYaml struct {
	Cluster clusterYaml `json:"cluster"`
}
type clusterYaml struct {
	CertificateAuthData string `yaml:"certificate-authority-data"`
	Server              string `yaml:"server"`
}

