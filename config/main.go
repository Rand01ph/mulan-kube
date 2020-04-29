package config

import (
	"bytes"
	"github.com/cuisongliu/drone-kube/tools"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"text/template"
)

var templateText = string(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: {{.KubeCa}}
    server: {{.KubeServer}}
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: {{.KubeAdmin}}
    client-key-data: {{.KubeAdminKey}}
`)

//var is global var
var (
	KubeCa     string
	KubeServer string

	KubeAdmin    string
	KubeAdminKey string
)

//Main is config command
func Main() {

	if KubeServer == "" {
		logger.Error("param server is null")
		return
	}
	if KubeCa == "" {
		logger.Error("param ca is null")
		return
	}
	if KubeAdmin == "" {
		logger.Error("param admin is null")
		return
	}
	if KubeAdminKey == "" {
		logger.Error("param admin key is null")
		return
	}
	var kubeconfig = tools.KubeConfigExists()
	var envMap = make(map[string]string, 4)
	envMap["KubeServer"] = KubeServer
	envMap["KubeCa"] = KubeCa
	envMap["KubeAdmin"] = KubeAdmin
	envMap["KubeAdminKey"] = KubeAdminKey
	tmpl, err := template.New("config").Parse(templateText)
	if err != nil {
		logger.Error("template parse failed:", err)
		return
	}
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	logger.Debug(&buffer)
	//write file
	_ = ioutil.WriteFile(kubeconfig, buffer.Bytes(), 0755)
}
