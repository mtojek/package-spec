package validator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/package-spec/code/go/internal/validator"
)

func TestValidate(t *testing.T) {
	pkgRootPath := "../../internal/validator/test/packages/good"
	errs := ValidateFromPath(pkgRootPath)
	require.NoError(t, errs)
}

func TestBadDeployVariants(t *testing.T) {
	pkgRootPath := "../../internal/validator/test/packages/bad_deploy_variants"
	errs := ValidateFromPath(pkgRootPath)
	require.Len(t, errs, 2)
	vErrs := errs.(interface{}).(validator.ValidationErrors)

	require.Equal(t,
		`file "../../internal/validator/test/packages/bad_deploy_variants/_dev/deploy/variants.yml" is invalid: field (root): default is required`,
		vErrs[0].Error())
	require.Equal(t,
		`file "../../internal/validator/test/packages/bad_deploy_variants/_dev/deploy/variants.yml" is invalid: field variants: Invalid type. Expected: object, given: array`,
		vErrs[1].Error())
}
