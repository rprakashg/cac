package errors

import "errors"

var (
	ErrResourceIsNil                       = errors.New("resource is nil")
	ErrResourceNameIsNil                   = errors.New("metadata.name is not set")
	ErrResourceOwnerIsNil                  = errors.New("metadata.owner not set")
	ErrResourceNotFound                    = errors.New("resource not found")
	ErrParentResourceNotFound              = errors.New("parent resource not found")
	ErrUpdatingResourceWithOwnerNotAllowed = errors.New("updating the resource is not allowed because it has an owner")
	ErrDeletingResourceWithOwnerNotAllowed = errors.New("deleting the resource is not allowed because it has an owner")
	ErrResourceNotEmpty                    = errors.New("resource cannot be deleted because it contains child resources")
	ErrDuplicateName                       = errors.New("a resource with this name already exists")
	ErrResourceVersionConflict             = errors.New("the object has been modified; please apply your changes to the latest version and try again")
	ErrIllegalResourceVersionFormat        = errors.New("resource version does not match the required integer format")
	ErrNoRowsUpdated                       = errors.New("no rows were updated; assuming resource version was updated or resource was deleted")
	ErrFieldSelectorSyntax                 = errors.New("invalid field selector syntax")
	ErrFieldSelectorParseFailed            = errors.New("failed to parse field selector")
	ErrFieldSelectorUnknownSelector        = errors.New("unknown or unsupported selector")
	ErrLabelSelectorSyntax                 = errors.New("invalid label selector syntax")
	ErrLabelSelectorParseFailed            = errors.New("failed to parse label selector")
	ErrAnnotationSelectorSyntax            = errors.New("invalid annotation selector syntax")
	ErrAnnotationSelectorParseFailed       = errors.New("failed to parse annotation selector")

	//policies
	ErrNoRenderedVersion = errors.New("no rendered version for policy")
)
