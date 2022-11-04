---
layout: "ibm"
page_title: "IBM : ibm_cm_object"
description: |-
  Get information about cm_object
subcategory: "Catalog Management API"
---

# ibm_cm_object

Provides a read-only data source for cm_object. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_object" "cm_object" {
	catalog_id = ibm_cm_object.cm_object.catalog_id
	object_id = ibm_cm_object.cm_object.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `catalog_id` - (Required, Forces new resource, String) Catalog identifier.
* `object_id` - (Required, Forces new resource, String) Object identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cm_object.
* `catalog_id` - (String) The id of the catalog containing this offering.

* `catalog_name` - (String) The name of the catalog.

* `catalog_object_id` - (String) unique id.

* `created` - (String) The date and time this catalog was created.

* `crn` - (String) The crn for this specific object.

* `data` - (Map) Map of data values for this object.

* `kind` - (String) Kind of object.

* `label` - (String) Display name in the requested language.

* `label_i18n` - (Map) A map of translated strings, by language code.

* `name` - (String) The programmatic name of this object.

* `parent_id` - (String) The parent for this specific object.

* `publish` - (List) Publish information.
Nested scheme for **publish**:
	* `ibm_approved` - (Boolean) Indicates if this offering has been approved for use by all IBMers.
	* `permit_ibm_public_publish` - (Boolean) Is it permitted to request publishing to IBM or Public.
	* `portal_approval_record` - (String) The portal's approval record ID.
	* `portal_url` - (String) The portal UI URL.
	* `public_approved` - (Boolean) Indicates if this offering has been approved for use by all IBM Cloud users.

* `rev` - (String) Cloudant revision.

* `short_description` - (String) Short description in the requested language.

* `short_description_i18n` - (Map) A map of translated strings, by language code.

* `state` - (List) Offering state.
Nested scheme for **state**:
	* `current` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `current_entered` - (String) Date and time of current request.
	* `pending` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `pending_requested` - (String) Date and time of pending request.
	* `previous` - (String) one of: new, validated, account-published, ibm-published, public-published.

* `tags` - (List) List of tags associated with this catalog.

* `updated` - (String) The date and time this catalog was last updated.

* `url` - (String) The url for this specific object.

