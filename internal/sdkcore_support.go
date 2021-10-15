package internal

type Registry = registry

func NewRegistry() *Registry { return newRegistry() }

func (r *Registry) GetWorkflowDefinition(wt WorkflowType) (WorkflowDefinition, error) {
	return r.getWorkflowDefinition(wt)
}
