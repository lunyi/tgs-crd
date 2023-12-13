package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// LobbySpec defines the desired state of Lobby
type LobbySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS -- desired state of cluster
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
	Domain   string `json:"domain"`
}

// LobbyStatus defines the observed state of Lobby.
// It should always be reconstructable from the state of the cluster and/or outside world.
type LobbyStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of cluster
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

// Lobby is the Schema for the lobbies API
// +k8s:openapi-gen=true
type Lobby struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LobbySpec   `json:"spec,omitempty"`
	Status LobbyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LobbyList contains a list of Lobby
type LobbyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lobby `json:"items"`
}
