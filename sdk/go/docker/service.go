// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package docker

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	// See Auth below for details.
	Auth ServiceAuthPtrOutput `pulumi:"auth"`
	// See Converge Config below for details.
	ConvergeConfig ServiceConvergeConfigPtrOutput `pulumi:"convergeConfig"`
	// See EndpointSpec below for details.
	EndpointSpec ServiceEndpointSpecOutput `pulumi:"endpointSpec"`
	// See Labels below for details.
	Labels ServiceLabelArrayOutput `pulumi:"labels"`
	// See Mode below for details.
	Mode ServiceModeOutput `pulumi:"mode"`
	// The name of the Docker service.
	Name pulumi.StringOutput `pulumi:"name"`
	// See RollbackConfig below for details.
	RollbackConfig ServiceRollbackConfigPtrOutput `pulumi:"rollbackConfig"`
	// See TaskSpec below for details.
	TaskSpec ServiceTaskSpecOutput `pulumi:"taskSpec"`
	// See UpdateConfig below for details.
	UpdateConfig ServiceUpdateConfigPtrOutput `pulumi:"updateConfig"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.TaskSpec == nil {
		return nil, errors.New("missing required argument 'TaskSpec'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("docker:index/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("docker:index/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// See Auth below for details.
	Auth *ServiceAuth `pulumi:"auth"`
	// See Converge Config below for details.
	ConvergeConfig *ServiceConvergeConfig `pulumi:"convergeConfig"`
	// See EndpointSpec below for details.
	EndpointSpec *ServiceEndpointSpec `pulumi:"endpointSpec"`
	// See Labels below for details.
	Labels []ServiceLabel `pulumi:"labels"`
	// See Mode below for details.
	Mode *ServiceMode `pulumi:"mode"`
	// The name of the Docker service.
	Name *string `pulumi:"name"`
	// See RollbackConfig below for details.
	RollbackConfig *ServiceRollbackConfig `pulumi:"rollbackConfig"`
	// See TaskSpec below for details.
	TaskSpec *ServiceTaskSpec `pulumi:"taskSpec"`
	// See UpdateConfig below for details.
	UpdateConfig *ServiceUpdateConfig `pulumi:"updateConfig"`
}

type ServiceState struct {
	// See Auth below for details.
	Auth ServiceAuthPtrInput
	// See Converge Config below for details.
	ConvergeConfig ServiceConvergeConfigPtrInput
	// See EndpointSpec below for details.
	EndpointSpec ServiceEndpointSpecPtrInput
	// See Labels below for details.
	Labels ServiceLabelArrayInput
	// See Mode below for details.
	Mode ServiceModePtrInput
	// The name of the Docker service.
	Name pulumi.StringPtrInput
	// See RollbackConfig below for details.
	RollbackConfig ServiceRollbackConfigPtrInput
	// See TaskSpec below for details.
	TaskSpec ServiceTaskSpecPtrInput
	// See UpdateConfig below for details.
	UpdateConfig ServiceUpdateConfigPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// See Auth below for details.
	Auth *ServiceAuth `pulumi:"auth"`
	// See Converge Config below for details.
	ConvergeConfig *ServiceConvergeConfig `pulumi:"convergeConfig"`
	// See EndpointSpec below for details.
	EndpointSpec *ServiceEndpointSpec `pulumi:"endpointSpec"`
	// See Labels below for details.
	Labels []ServiceLabel `pulumi:"labels"`
	// See Mode below for details.
	Mode *ServiceMode `pulumi:"mode"`
	// The name of the Docker service.
	Name *string `pulumi:"name"`
	// See RollbackConfig below for details.
	RollbackConfig *ServiceRollbackConfig `pulumi:"rollbackConfig"`
	// See TaskSpec below for details.
	TaskSpec ServiceTaskSpec `pulumi:"taskSpec"`
	// See UpdateConfig below for details.
	UpdateConfig *ServiceUpdateConfig `pulumi:"updateConfig"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// See Auth below for details.
	Auth ServiceAuthPtrInput
	// See Converge Config below for details.
	ConvergeConfig ServiceConvergeConfigPtrInput
	// See EndpointSpec below for details.
	EndpointSpec ServiceEndpointSpecPtrInput
	// See Labels below for details.
	Labels ServiceLabelArrayInput
	// See Mode below for details.
	Mode ServiceModePtrInput
	// The name of the Docker service.
	Name pulumi.StringPtrInput
	// See RollbackConfig below for details.
	RollbackConfig ServiceRollbackConfigPtrInput
	// See TaskSpec below for details.
	TaskSpec ServiceTaskSpecInput
	// See UpdateConfig below for details.
	UpdateConfig ServiceUpdateConfigPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
