# API Reference



# PhonebookService


## CreateContact

> **rpc** CreateContact([CreateContactRequest](#createcontactrequest))
    [CreateContactResponse](#createcontactresponse)


## GetContact

> **rpc** GetContact([GetContactRequest](#getcontactrequest))
    [GetContactResponse](#getcontactresponse)


## ListContacts

> **rpc** ListContacts([ListContactsRequest](#listcontactsrequest))
    [ListContactsResponse](#listcontactsresponse)


## UpdateContact

> **rpc** UpdateContact([UpdateContactRequest](#updatecontactrequest))
    [UpdateContactResponse](#updatecontactresponse)


## DeleteContact

> **rpc** DeleteContact([DeleteContactRequest](#deletecontactrequest))
    [DeleteContactResponse](#deletecontactresponse)


 <!-- end methods -->
 <!-- end services -->

# Messages


## Contact



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | none |
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
