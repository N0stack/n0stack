package target

import "github.com/n0stack/n0core/model"

// Target is application service to apply resources with some framework like KVM and iproute2.
//
// A target manage only one type `*/*/*` of resource like `resource/network/flat`.
// Directory structure and class name is ruled by resource type.
// For example, `resource/network/flat` define `class Flat` which is placed on `n0core.resource.network.flat`.
//
// Do not implement that resource is kill when target is killed.
//
// Args:
// 	support_model: Model type which is supported on each target.
//
// Example:
//
// 	in `n0core.target.vm.example`
// 	>>> class Exapmle(Target):
// 	>>>     def __init__(self):
// 	>>>         super().__init__("resource/vm/example")
//
// Dependency packages:
//
// Orchestration pipeline:
//
type (
	Target interface {
		// Operations return slice of functions as operations on the task, when task can be processed on the state.
		Operations(state, task string) ([]*Operation, error)

		// ManagingType return supporting model type.
		ManagingType() string

		// Initialize initialize target to orchestrate resources.
		//
		// Ex. detect already orchestrated resource and test automatically.
		// Initialize() (string, bool)

		// Do test in Initialize to check whether target can orchestrate resource rightly.
		// test() (string, bool)

		ParseModel(m model.AbstractModel) (bool, string)
		CheckState(m model.AbstractModel) (bool, string)
	}

	Operation struct {
		Function func(n model.AbstractModel) (bool, string)
		Name     string
	}
)
