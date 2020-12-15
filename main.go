package main

import (
	"fmt"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	virtcli "kubevirt.io/client-go/kubecli"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	// kubecli.DefaultClientConfig() prepares config using kubeconfig.
	// typically, you need to set env variable, KUBECONFIG=<path-to-kubeconfig>/.kubeconfig
	clientConfig := virtcli.DefaultClientConfig(&pflag.FlagSet{})
	// retrive default namespace.


	// get the kubevirt client, using which kubevirt resources can be managed.
	virtClient, err := virtcli.GetKubevirtClientFromClientConfig(clientConfig)
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt client: %v\n", err)
	}

	// Fetch list of VMs & VMIs
	vmList, err := virtClient.VirtualMachine("default").List(&metav1.ListOptions{})
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt vm list: %v\n", err)
	}

	vmiList, err := virtClient.VirtualMachineInstance("default").List(&metav1.ListOptions{})
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt vmi list: %v\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	fmt.Fprintln(w, "Type\tName\tNamespace\tStatus")

	for _, vm := range vmList.Items {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n", vm.Kind, vm.Name, vm.Namespace, vm.Status.Ready)
		if !vm.Status.Ready {
			virtClient.VirtualMachine("default").Stop(vm.Name)
		}
	}
	for _, vmi := range vmiList.Items {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n", vmi.Kind, vmi.Name, vmi.Namespace, vmi.Status.Phase)
	}
	w.Flush()
}