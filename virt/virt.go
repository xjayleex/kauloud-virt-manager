package virt

import (
	v1 "kubevirt.io/client-go/api/v1"
	"kubevirt.io/client-go/kubecli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KubeVirtManager struct {
	virtClient kubecli.KubevirtClient
}

func NewKubeVirtManager () *KubeVirtManager{
	// TODO gRPC Wrapping
	return &KubeVirtManager{}
}

func (o *KubeVirtManager) VirtClient() kubecli.KubevirtClient {
	return o.virtClient
}

func (o *KubeVirtManager) ListVM (namespace string, options *metav1.ListOptions) (*v1.VirtualMachineList, error) {
	vmList, err := o.VirtClient().VirtualMachine(namespace).List(options)
	return vmList, err
}

func (o *KubeVirtManager) CreateVM (vm *v1.VirtualMachine, namespace string) (err error) {
	o.VirtClient().VirtualMachine(namespace).Create(vm)
	return err
}

func (o *KubeVirtManager) DeleteVM (vm *v1.VirtualMachine, options *metav1.DeleteOptions, namespace string) (err error) {
	o.VirtClient().VirtualMachine(namespace).Delete(vm.Name, options)
}

func (o *KubeVirtManager) StartVM(vm *v1.VirtualMachine, namespace string) (err error){
	// check vm exists
	err = o.VirtClient().VirtualMachine(namespace).Start(vm.Name)
	return err
}

func (o *KubeVirtManager) StopVM(vm *v1.VirtualMachine, namespace string) (err error) {
	err = o.VirtClient().VirtualMachine(namespace).Stop(vm.Name)
	return err
}

func (o *KubeVirtManager) RestartVM(vm *v1.VirtualMachine, namespace string) (err error) {
	// check vmi running
	err = o.VirtClient().VirtualMachine(namespace).Restart(vm.Name)
	return err
}