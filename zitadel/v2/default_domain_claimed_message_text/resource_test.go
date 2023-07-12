package default_domain_claimed_message_text_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/admin"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
)

func TestAccZITADELDefaultDomainClaimedMessageTextSameLanguage(t *testing.T) {
	resourceName := "zitadel_default_domain_claimed_message_text"
	initialProperty := "initialtitle"
	updatedProperty := "updatedtitle"
	language := "en"
	frame, err := test_utils.NewInstanceTestFrame(resourceName)
	if err != nil {
		t.Fatalf("setting up test context failed: %v", err)
	}
	test_utils.RunLifecyleTest(
		t,
		frame.BaseTestFrame,
		func(configProperty, _ string) string {
			return fmt.Sprintf(`
resource "%s" "%s" {
  language    = "%s"
  title       = "%s"
  pre_header  = "pre_header example"
  subject     = "subject example"
  greeting    = "greeting example"
  text        = "text example"
  button_text = "button_text example"
  footer_text = "footer_text example"
}`, resourceName, frame.UniqueResourcesID, language, configProperty)
		},
		initialProperty, updatedProperty,
		"", "",
		checkRemoteProperty(frame, language),
		regexp.MustCompile(`^en$`),
		// When deleted, the default should be returned
		checkRemoteProperty(frame, language)("ZITADEL - Domain has been claimed"),
		nil, nil, "", "",
	)
}

func checkRemoteProperty(frame *test_utils.InstanceTestFrame, lang string) func(string) resource.TestCheckFunc {
	return func(expect string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			remoteResource, err := frame.GetCustomDomainClaimedMessageText(frame, &admin.GetCustomDomainClaimedMessageTextRequest{Language: lang})
			if err != nil {
				return err
			}
			actual := remoteResource.GetCustomText().GetTitle()
			if actual != expect {
				return fmt.Errorf("expected %s, actual: %s", expect, actual)
			}
			return nil
		}
	}
}
