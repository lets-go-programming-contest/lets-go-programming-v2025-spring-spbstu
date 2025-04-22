# API Reference



# ContactManagerService
Service for managing contacts

## CreateContact

> **rpc** CreateContact([CreateContactRequest](#createcontactrequest))
    [CreateContactResponse](#createcontactresponse)

Create a new contact
## GetContact

> **rpc** GetContact([GetContactRequest](#getcontactrequest))
    [GetContactResponse](#getcontactresponse)

Get a contact by ID
## ListContacts

> **rpc** ListContacts([ListContactsRequest](#listcontactsrequest))
    [ListContactsResponse](#listcontactsresponse)

List all contacts
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


## Contact



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | none |
| name | [ string](#string) | none |
| phone | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## CreateContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | none |
| phone | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## CreateContactResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| contact | [ Contact](#contact) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## DeleteContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## DeleteContactResponse


 <!-- end HasFields -->


## GetContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GetContactResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| contact | [ Contact](#contact) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ListContactsRequest


 <!-- end HasFields -->


## ListContactsResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| contacts | [repeated Contact](#contact) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## UpdateContactRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | none |
| name | [ string](#string) | none |
| phone | [ string](#string) | none |
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
