# API Reference



# ContactManagerService
Service for managing contacts

## AddContact

> **rpc** AddContact([AddContactRequest](#addcontactrequest))
    [AddContactResponse](#addcontactresponse)

Add a new contact
## GetContact

> **rpc** GetContact([GetContactRequest](#getcontactrequest))
    [GetContactResponse](#getcontactresponse)

Get a contact by ID
## GetAllContacts

> **rpc** GetAllContacts([GetAllContactsRequest](#getallcontactsrequest))
    [GetAllContactsResponse](#getallcontactsresponse)

Get all contacts
## UpdateContact

> **rpc** UpdateContact([UpdateContactRequest](#updatecontactrequest))
    [UpdateContactResponse](#updatecontactresponse)

Update a contact
## DeleteContact

> **rpc** DeleteContact([DeleteContactRequest](#deletecontactrequest))
    [DeleteContactResponse](#deletecontactresponse)

Delete a contact
 <!-- end methods -->
 <!-- end services -->

# Messages


## AddContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name contact is associated with |
| number | [ string](#string) | Phone number contact is associated with |
 <!-- end Fields -->
 <!-- end HasFields -->


## AddContactResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| contact | [ Contact](#contact) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Contact



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id of contact's record in DB |
| name | [ string](#string) | Name contact is associated with |
| number | [ string](#string) | Phone number contact is associated with |
 <!-- end Fields -->
 <!-- end HasFields -->


## DeleteContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id of contact's record in DB |
 <!-- end Fields -->
 <!-- end HasFields -->


## DeleteContactResponse


 <!-- end HasFields -->


## GetAllContactsRequest


 <!-- end HasFields -->


## GetAllContactsResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| contacts | [repeated Contact](#contact) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GetContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id of contact's record in DB |
 <!-- end Fields -->
 <!-- end HasFields -->


## GetContactResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| contact | [ Contact](#contact) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## UpdateContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id of contact's record in DB |
| name | [ string](#string) | Name contact is associated with |
| number | [ string](#string) | Phone number contact is associated with |
 <!-- end Fields -->
 <!-- end HasFields -->


## UpdateContactResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| contact | [ Contact](#contact) | none |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->
 <!-- end Files -->
