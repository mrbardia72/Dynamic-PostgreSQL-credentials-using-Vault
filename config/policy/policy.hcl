# Read-only policy --> sample test policy config file HCL
path "secretv1/constrained/*" {
  capabilities = ["read"]
}