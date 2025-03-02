---
layout: 'okta'
page_title: 'Okta: okta_app_secure_password_store'
sidebar_current: 'docs-okta-resource-app-secure-password-store'
description: |-
  Creates a Secure Password Store Application.
---

# okta_app_secure_password_store

This resource allows you to create and configure a Secure Password Store Application.

## Example Usage

```hcl
resource "okta_app_secure_password_store" "example" {
  label              = "example"
  username_field     = "user"
  password_field     = "pass"
  url                = "https://test.com"
  credentials_scheme = "ADMIN_SETS_CREDENTIALS"
}
```

## Argument Reference

The following arguments are supported:

- `accessibility_error_redirect_url` - (Optional) Custom error page URL.

- `accessibility_login_redirect_url` - (Optional) Custom login page for this application.

- `accessibility_self_service` - (Optional) Enable self-service. By default, it is `false`.

- `admin_note` - (Optional) Application notes for admins.

- `app_links_json` - (Optional) Displays specific appLinks for the app. The value for each application link should be boolean.

- `auto_submit_toolbar` - (Optional) Display auto submit toolbar.

- `credentials_scheme` - (Optional) Application credentials scheme. Can be set to `"EDIT_USERNAME_AND_PASSWORD"`, `"ADMIN_SETS_CREDENTIALS"`, `"EDIT_PASSWORD_ONLY"`, `"EXTERNAL_PASSWORD_SYNC"`, or `"SHARED_USERNAME_AND_PASSWORD"`.

- `enduser_note` - (Optional) Application notes for end users.

- `groups` - (Optional) Groups associated with the application. See `okta_app_group_assignment` for a more flexible approach.
  - `DEPRECATED`: Please replace usage with the `okta_app_group_assignments` (or `okta_app_group_assignment`) resource.

- `hide_ios` - (Optional) Do not display application icon on mobile app.

- `hide_web` - (Optional) Do not display application icon to users.

- `label` - (Required) The display name of the Application.

- `logo` - (Optional) Local file path to the logo. The file must be in PNG, JPG, or GIF format, and less than 1 MB in size.

- `optional_field1` - (Optional) Name of optional param in the login form.

- `optional_field1_value` - (Optional) Name of optional value in the login form.

- `optional_field2` - (Optional) Name of optional param in the login form.

- `optional_field2_value` - (Optional) Name of optional value in the login form.

- `optional_field3` - (Optional) Name of optional param in the login form.

- `optional_field3_value` - (Optional) Name of optional value in the login form.

- `password_field` - (Required) Login password field.

- `reveal_password` - (Optional) Allow user to reveal password. It can not be set to `true` if `credentials_scheme` is `"ADMIN_SETS_CREDENTIALS"`, `"SHARED_USERNAME_AND_PASSWORD"` or `"EXTERNAL_PASSWORD_SYNC"`.

- `shared_password` - (Optional) Shared password, required for certain schemes.

- `shared_username` - (Optional) Shared username, required for certain schemes.

- `skip_groups` - (Optional) Indicator that allows the app to skip `groups` sync (it's also can be provided during import). Default is `false`.

- `skip_users` - (Optional) Indicator that allows the app to skip `users` sync (it's also can be provided during import). Default is `false`.

- `status` - (Optional) Status of application. By default, it is `"ACTIVE"`.

- `url` - (Required) Login URL.

- `user_name_template` - (Optional) Username template. Default: `"${source.login}"`

- `user_name_template_push_status` - (Optional) Push username on update. Valid values: `"PUSH"` and `"DONT_PUSH"`.

- `user_name_template_suffix` - (Optional) Username template suffix.

- `user_name_template_type` - (Optional) Username template type. Default: `"BUILT_IN"`.

- `username_field` - (Required) Login username field.

- `users` - (Optional) The users assigned to the application. See `okta_app_user` for a more flexible approach.
  - `DEPRECATED`: Please replace usage with the `okta_app_user` resource.

## Attributes Reference

- `name` - Name assigned to the application by Okta.

- `sign_on_mode` - Sign-on mode of application.

## Timeouts

The `timeouts` block allows you to specify custom [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions: 

- `create` - Create timeout if syncing users/groups (default 1 hour).

- `update` - Update timeout if syncing users/groups (default 1 hour).

- `read` - Read timeout if syncing users/groups (default 1 hour).

## Import

Secure Password Store Application can be imported via the Okta ID.

```
$ terraform import okta_app_secure_password_store.example &#60;app id&#62;
```

It's also possible to import app without groups or/and users. In this case ID may look like this:

```
$ terraform import okta_app_basic_auth.example &#60;app id&#62;/skip_users

$ terraform import okta_app_basic_auth.example &#60;app id&#62;/skip_users/skip_groups

$ terraform import okta_app_basic_auth.example &#60;app id&#62;/skip_groups
```
