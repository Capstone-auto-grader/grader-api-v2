// Package graderd implements a grader API server based on the protobuf definition
// in graderpb/grader.proto.
//
// Usage
//	srv := NewGraderService(schr, db, webAddr)
//
// Note that the API is not yet stable, please use with caution.
package graderd
