/* Created by "go tool cgo" - DO NOT EDIT. */

/* package _/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind */

/* Start of preamble from import "C" comments.  */


#line 5 "/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/classes.go"

#include <stdlib.h> // for free()
#include <jni.h>
#include "seq.h"
#include "classes.h"

#line 1 "cgo-generated-wrapper"

#line 7 "/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go"

#include <stdlib.h>
#include <stdint.h>
#include "seq.h"
#include "holochain.h"


#line 1 "cgo-generated-wrapper"

#line 7 "/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_main.go"

#include <stdlib.h>
#include <stdint.h>
#include "seq.h"
#include "universe.h"


#line 1 "cgo-generated-wrapper"

#line 11 "/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq.go"
 #include <stdlib.h>
 #include "seq.h"

#line 1 "cgo-generated-wrapper"

#line 11 "/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go"


#include <jni.h>
#include <stdint.h>
#include <stdlib.h>
#include "seq.h"

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt32 GoInt;
typedef GoUint32 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_32_bit_pointer_matching_GoInt[sizeof(void*)==32/8 ? 1:-1];

typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


extern void initClasses();

extern nstring proxyholochain_ActionBridge_Name(int32_t p0);

extern int32_t new_holochain_ActionBridge();

extern nstring proxyholochain_ActionCall_Name(int32_t p0);

extern int32_t new_holochain_ActionCall();

extern int32_t proxyholochain_ActionCommit_CheckValidationRequest(int32_t p0, int32_t p1);

extern int32_t proxyholochain_ActionCommit_Entry(int32_t p0);

extern nstring proxyholochain_ActionCommit_EntryType(int32_t p0);

extern nstring proxyholochain_ActionCommit_Name(int32_t p0);

extern void proxyholochain_ActionCommit_SetHeader(int32_t p0, int32_t p1);

extern int32_t new_holochain_ActionCommit();

extern nstring proxyholochain_ActionDebug_Name(int32_t p0);

extern int32_t new_holochain_ActionDebug();

extern int32_t proxyholochain_ActionDel_CheckValidationRequest(int32_t p0, int32_t p1);

extern int32_t proxyholochain_ActionDel_Entry(int32_t p0);

extern nstring proxyholochain_ActionDel_EntryType(int32_t p0);

extern nstring proxyholochain_ActionDel_Name(int32_t p0);

extern void proxyholochain_ActionDel_SetHeader(int32_t p0, int32_t p1);

extern int32_t new_holochain_ActionDel();

extern nstring proxyholochain_ActionGet_Name(int32_t p0);

extern int32_t new_holochain_ActionGet();

extern nstring proxyholochain_ActionGetBridges_Name(int32_t p0);

extern int32_t new_holochain_ActionGetBridges();

extern nstring proxyholochain_ActionGetLinks_Name(int32_t p0);

extern int32_t new_holochain_ActionGetLinks();

extern int32_t proxyholochain_ActionLink_CheckValidationRequest(int32_t p0, int32_t p1);

extern nstring proxyholochain_ActionLink_Name(int32_t p0);

extern int32_t new_holochain_ActionLink();

extern nstring proxyholochain_ActionListAdd_Name(int32_t p0);

extern int32_t new_holochain_ActionListAdd();

extern nstring proxyholochain_ActionMakeHash_Name(int32_t p0);

extern int32_t new_holochain_ActionMakeHash();

extern int32_t proxyholochain_ActionMod_CheckValidationRequest(int32_t p0, int32_t p1);

extern int32_t proxyholochain_ActionMod_Entry(int32_t p0);

extern nstring proxyholochain_ActionMod_EntryType(int32_t p0);

extern nstring proxyholochain_ActionMod_Name(int32_t p0);

extern void proxyholochain_ActionMod_SetHeader(int32_t p0, int32_t p1);

extern int32_t new_holochain_ActionMod();

extern void proxyholochain_ActionModAgent_Revocation_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ActionModAgent_Revocation_Get(int32_t p0);

extern nstring proxyholochain_ActionModAgent_Name(int32_t p0);

extern int32_t new_holochain_ActionModAgent();

extern nstring proxyholochain_ActionProperty_Name(int32_t p0);

extern int32_t new_holochain_ActionProperty();

extern int32_t proxyholochain_ActionPut_CheckValidationRequest(int32_t p0, int32_t p1);

extern nstring proxyholochain_ActionPut_Name(int32_t p0);

extern int32_t new_holochain_ActionPut();

extern nstring proxyholochain_ActionQuery_Name(int32_t p0);

extern int32_t new_holochain_ActionQuery();

extern nstring proxyholochain_ActionSend_Name(int32_t p0);

extern int32_t new_holochain_ActionSend();

extern nstring proxyholochain_ActionSign_Name(int32_t p0);

extern int32_t new_holochain_ActionSign();

extern nstring proxyholochain_ActionVerifySignature_Name(int32_t p0);

extern int32_t new_holochain_ActionVerifySignature();

extern void proxyholochain_AgentEntry_Revocation_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_AgentEntry_Revocation_Get(int32_t p0);

extern void proxyholochain_AgentEntry_PublicKey_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_AgentEntry_PublicKey_Get(int32_t p0);

extern int32_t new_holochain_AgentEntry();

extern void proxyholochain_AgentFixture_Hash_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AgentFixture_Hash_Get(int32_t p0);

extern void proxyholochain_AgentFixture_Identity_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AgentFixture_Identity_Get(int32_t p0);

extern int32_t new_holochain_AgentFixture();

extern void proxyholochain_AppMsg_ZomeType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppMsg_ZomeType_Get(int32_t p0);

extern void proxyholochain_AppMsg_Body_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppMsg_Body_Get(int32_t p0);

extern int32_t new_holochain_AppMsg();

extern void proxyholochain_AppPackage_Version_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppPackage_Version_Get(int32_t p0);

extern void proxyholochain_AppPackage_Generator_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppPackage_Generator_Get(int32_t p0);

extern int32_t new_holochain_AppPackage();

extern void proxyholochain_AppPackageScenario_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppPackageScenario_Name_Get(int32_t p0);

extern int32_t new_holochain_AppPackageScenario();

extern void proxyholochain_AppPackageTests_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppPackageTests_Name_Get(int32_t p0);

extern int32_t new_holochain_AppPackageTests();

extern void proxyholochain_AppPackageUIFile_FileName_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppPackageUIFile_FileName_Get(int32_t p0);

extern void proxyholochain_AppPackageUIFile_Data_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppPackageUIFile_Data_Get(int32_t p0);

extern void proxyholochain_AppPackageUIFile_Encoding_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_AppPackageUIFile_Encoding_Get(int32_t p0);

extern int32_t new_holochain_AppPackageUIFile();

extern void proxyholochain_Arg_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Arg_Name_Get(int32_t p0);

extern void proxyholochain_Arg_Optional_Set(int32_t p0, char p1);

extern char proxyholochain_Arg_Optional_Get(int32_t p0);

extern int32_t new_holochain_Arg();

extern void proxyholochain_BSReq_Version_Set(int32_t p0, nint p1);

extern nint proxyholochain_BSReq_Version_Get(int32_t p0);

extern void proxyholochain_BSReq_NodeID_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_BSReq_NodeID_Get(int32_t p0);

extern void proxyholochain_BSReq_NodeAddr_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_BSReq_NodeAddr_Get(int32_t p0);

extern int32_t new_holochain_BSReq();

extern void proxyholochain_BSResp_Remote_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_BSResp_Remote_Get(int32_t p0);

extern int32_t new_holochain_BSResp();

extern void proxyholochain_Bridge_Token_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Bridge_Token_Get(int32_t p0);

extern void proxyholochain_Bridge_Side_Set(int32_t p0, nint p1);

extern nint proxyholochain_Bridge_Side_Get(int32_t p0);

extern int32_t new_holochain_Bridge();

extern void proxyholochain_BridgeApp_H_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_BridgeApp_H_Get(int32_t p0);

extern void proxyholochain_BridgeApp_Side_Set(int32_t p0, nint p1);

extern nint proxyholochain_BridgeApp_Side_Get(int32_t p0);

extern void proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get(int32_t p0);

extern void proxyholochain_BridgeApp_BridgeGenesisDataTo_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_BridgeApp_BridgeGenesisDataTo_Get(int32_t p0);

extern void proxyholochain_BridgeApp_Port_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_BridgeApp_Port_Get(int32_t p0);

extern int32_t new_holochain_BridgeApp();

extern nint proxyholochain_Bucket_Len(int32_t p0);

extern int32_t new_holochain_Bucket();

extern void proxyholochain_BytesSent_Bytes_Set(int32_t p0, int64_t p1);

extern int64_t proxyholochain_BytesSent_Bytes_Get(int32_t p0);

extern int32_t new_holochain_BytesSent();

extern void proxyholochain_Callback_Function_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Callback_Function_Get(int32_t p0);

extern void proxyholochain_Callback_ID_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Callback_ID_Get(int32_t p0);

extern int32_t new_holochain_Callback();

extern void proxyholochain_Capability_Token_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Capability_Token_Get(int32_t p0);

extern int32_t new_holochain_Capability();

extern void proxyholochain_Chain_Close(int32_t p0);

/* Return type for proxyholochain_Chain_JSON */
struct proxyholochain_Chain_JSON_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_Chain_JSON_return proxyholochain_Chain_JSON(int32_t p0);

extern nint proxyholochain_Chain_Length(int32_t p0);

extern int32_t proxyholochain_Chain_Nth(int32_t p0, nint p1);

extern nstring proxyholochain_Chain_String(int32_t p0);

extern int32_t proxyholochain_Chain_Top(int32_t p0);

extern int32_t proxyholochain_Chain_Validate(int32_t p0, char p1);

extern int32_t new_holochain_Chain();

extern void proxyholochain_ChainPair_Header_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_ChainPair_Header_Get(int32_t p0);

extern void proxyholochain_ChainPair_Entry_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_ChainPair_Entry_Get(int32_t p0);

extern int32_t new_holochain_ChainPair();

extern void proxyholochain_Change_Message_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Change_Message_Get(int32_t p0);

extern void proxyholochain_Change_AsOf_Set(int32_t p0, nint p1);

extern nint proxyholochain_Change_AsOf_Get(int32_t p0);

extern void proxyholochain_Change_Log(int32_t p0);

extern int32_t new_holochain_Change();

extern void proxyholochain_CloneSpec_Role_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_CloneSpec_Role_Get(int32_t p0);

extern void proxyholochain_CloneSpec_Number_Set(int32_t p0, nint p1);

extern nint proxyholochain_CloneSpec_Number_Get(int32_t p0);

extern int32_t new_holochain_CloneSpec();

extern int32_t new_holochain_CloserPeersResp();

extern void proxyholochain_Config_Port_Set(int32_t p0, nint p1);

extern nint proxyholochain_Config_Port_Get(int32_t p0);

extern void proxyholochain_Config_EnableMDNS_Set(int32_t p0, char p1);

extern char proxyholochain_Config_EnableMDNS_Get(int32_t p0);

extern void proxyholochain_Config_PeerModeAuthor_Set(int32_t p0, char p1);

extern char proxyholochain_Config_PeerModeAuthor_Get(int32_t p0);

extern void proxyholochain_Config_PeerModeDHTNode_Set(int32_t p0, char p1);

extern char proxyholochain_Config_PeerModeDHTNode_Get(int32_t p0);

extern void proxyholochain_Config_EnableNATUPnP_Set(int32_t p0, char p1);

extern char proxyholochain_Config_EnableNATUPnP_Get(int32_t p0);

extern void proxyholochain_Config_BootstrapServer_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Config_BootstrapServer_Get(int32_t p0);

extern int32_t proxyholochain_Config_Setup(int32_t p0);

extern int32_t proxyholochain_Config_SetupLogging(int32_t p0);

extern int32_t new_holochain_Config();

extern void proxyholochain_DHT_Close(int32_t p0);

/* Return type for proxyholochain_DHT_DumpIdx */
struct proxyholochain_DHT_DumpIdx_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_DHT_DumpIdx_return proxyholochain_DHT_DumpIdx(int32_t p0, nint p1);

/* Return type for proxyholochain_DHT_DumpIdxJSON */
struct proxyholochain_DHT_DumpIdxJSON_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_DHT_DumpIdxJSON_return proxyholochain_DHT_DumpIdxJSON(int32_t p0, nint p1);

/* Return type for proxyholochain_DHT_GetIdx */
struct proxyholochain_DHT_GetIdx_return {
	nint r0;
	int32_t r1;
};

extern struct proxyholochain_DHT_GetIdx_return proxyholochain_DHT_GetIdx(int32_t p0);

extern int32_t proxyholochain_DHT_HandleGossipPuts(int32_t p0);

extern int32_t proxyholochain_DHT_HandleGossipWiths(int32_t p0);

/* Return type for proxyholochain_DHT_JSON */
struct proxyholochain_DHT_JSON_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_DHT_JSON_return proxyholochain_DHT_JSON(int32_t p0);

extern int32_t proxyholochain_DHT_SetupDHT(int32_t p0);

extern int32_t proxyholochain_DHT_Start(int32_t p0);

extern nstring proxyholochain_DHT_String(int32_t p0);

extern int32_t new_holochain_DHT();

extern void proxyholochain_DHTConfig_HashType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_DHTConfig_HashType_Get(int32_t p0);

extern void proxyholochain_DHTConfig_NeighborhoodSize_Set(int32_t p0, nint p1);

extern nint proxyholochain_DHTConfig_NeighborhoodSize_Get(int32_t p0);

extern int32_t new_holochain_DHTConfig();

extern void proxyholochain_DNA_Version_Set(int32_t p0, nint p1);

extern nint proxyholochain_DNA_Version_Get(int32_t p0);

extern void proxyholochain_DNA_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_DNA_Name_Get(int32_t p0);

extern void proxyholochain_DNA_PropertiesSchema_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_DNA_PropertiesSchema_Get(int32_t p0);

extern void proxyholochain_DNA_AgentIdentitySchema_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_DNA_AgentIdentitySchema_Get(int32_t p0);

extern void proxyholochain_DNA_RequiresVersion_Set(int32_t p0, nint p1);

extern nint proxyholochain_DNA_RequiresVersion_Get(int32_t p0);

extern int32_t proxyholochain_DNA_NewUUID(int32_t p0);

extern int32_t new_holochain_DNA();

extern void proxyholochain_DNAFile_Version_Set(int32_t p0, nint p1);

extern nint proxyholochain_DNAFile_Version_Get(int32_t p0);

extern void proxyholochain_DNAFile_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_DNAFile_Name_Get(int32_t p0);

extern void proxyholochain_DNAFile_PropertiesSchemaFile_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_DNAFile_PropertiesSchemaFile_Get(int32_t p0);

extern void proxyholochain_DNAFile_RequiresVersion_Set(int32_t p0, nint p1);

extern nint proxyholochain_DNAFile_RequiresVersion_Get(int32_t p0);

extern int32_t new_holochain_DNAFile();

extern void proxyholochain_DelEntry_Message_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_DelEntry_Message_Get(int32_t p0);

extern int32_t new_holochain_DelEntry();

extern int32_t new_holochain_DelReq();

extern void proxyholochain_EntryDef_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDef_Name_Get(int32_t p0);

extern void proxyholochain_EntryDef_DataFormat_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDef_DataFormat_Get(int32_t p0);

extern void proxyholochain_EntryDef_Sharing_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDef_Sharing_Get(int32_t p0);

extern void proxyholochain_EntryDef_Schema_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDef_Schema_Get(int32_t p0);

extern int32_t proxyholochain_EntryDef_BuildJSONSchemaValidator(int32_t p0, nstring p1);

extern int32_t proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString(int32_t p0, nstring p1);

extern char proxyholochain_EntryDef_IsSysEntry(int32_t p0);

extern char proxyholochain_EntryDef_IsVirtualEntry(int32_t p0);

extern int32_t new_holochain_EntryDef();

extern void proxyholochain_EntryDefFile_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDefFile_Name_Get(int32_t p0);

extern void proxyholochain_EntryDefFile_DataFormat_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDefFile_DataFormat_Get(int32_t p0);

extern void proxyholochain_EntryDefFile_Schema_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDefFile_Schema_Get(int32_t p0);

extern void proxyholochain_EntryDefFile_SchemaFile_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDefFile_SchemaFile_Get(int32_t p0);

extern void proxyholochain_EntryDefFile_Sharing_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_EntryDefFile_Sharing_Get(int32_t p0);

extern int32_t new_holochain_EntryDefFile();

extern void proxyholochain_ErrorResponse_Code_Set(int32_t p0, nint p1);

extern nint proxyholochain_ErrorResponse_Code_Get(int32_t p0);

extern void proxyholochain_ErrorResponse_Message_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ErrorResponse_Message_Get(int32_t p0);

extern int32_t proxyholochain_ErrorResponse_DecodeResponseError(int32_t p0);

extern int32_t new_holochain_ErrorResponse();

extern int32_t new_holochain_FindNodeReq();

extern void proxyholochain_FunctionDef_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_FunctionDef_Name_Get(int32_t p0);

extern void proxyholochain_FunctionDef_CallingType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_FunctionDef_CallingType_Get(int32_t p0);

extern void proxyholochain_FunctionDef_Exposure_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_FunctionDef_Exposure_Get(int32_t p0);

extern char proxyholochain_FunctionDef_ValidExposure(int32_t p0, nstring p1);

extern int32_t new_holochain_FunctionDef();

extern void proxyholochain_GetLinksOptions_Load_Set(int32_t p0, char p1);

extern char proxyholochain_GetLinksOptions_Load_Get(int32_t p0);

extern void proxyholochain_GetLinksOptions_StatusMask_Set(int32_t p0, nint p1);

extern nint proxyholochain_GetLinksOptions_StatusMask_Get(int32_t p0);

extern int32_t new_holochain_GetLinksOptions();

extern void proxyholochain_GetOptions_StatusMask_Set(int32_t p0, nint p1);

extern nint proxyholochain_GetOptions_StatusMask_Get(int32_t p0);

extern void proxyholochain_GetOptions_GetMask_Set(int32_t p0, nint p1);

extern nint proxyholochain_GetOptions_GetMask_Get(int32_t p0);

extern void proxyholochain_GetOptions_Local_Set(int32_t p0, char p1);

extern char proxyholochain_GetOptions_Local_Get(int32_t p0);

extern int32_t new_holochain_GetOptions();

extern void proxyholochain_GetReq_StatusMask_Set(int32_t p0, nint p1);

extern nint proxyholochain_GetReq_StatusMask_Get(int32_t p0);

extern void proxyholochain_GetReq_GetMask_Set(int32_t p0, nint p1);

extern nint proxyholochain_GetReq_GetMask_Get(int32_t p0);

extern int32_t new_holochain_GetReq();

extern void proxyholochain_GetResp_EntryType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_GetResp_EntryType_Get(int32_t p0);

extern void proxyholochain_GetResp_FollowHash_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_GetResp_FollowHash_Get(int32_t p0);

extern int32_t new_holochain_GetResp();

/* Return type for proxyholochain_GobEntry_Marshal */
struct proxyholochain_GobEntry_Marshal_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_GobEntry_Marshal_return proxyholochain_GobEntry_Marshal(int32_t p0);

extern int32_t proxyholochain_GobEntry_Unmarshal(int32_t p0, nbyteslice p1);

extern int32_t new_holochain_GobEntry();

extern int32_t new_holochain_Gossip();

extern void proxyholochain_GossipReq_MyIdx_Set(int32_t p0, nint p1);

extern nint proxyholochain_GossipReq_MyIdx_Get(int32_t p0);

extern void proxyholochain_GossipReq_YourIdx_Set(int32_t p0, nint p1);

extern nint proxyholochain_GossipReq_YourIdx_Get(int32_t p0);

extern int32_t new_holochain_GossipReq();

extern void proxyholochain_Header_Type_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Header_Type_Get(int32_t p0);

/* Return type for proxyholochain_Header_Marshal */
struct proxyholochain_Header_Marshal_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_Header_Marshal_return proxyholochain_Header_Marshal(int32_t p0);

extern int32_t proxyholochain_Header_Unmarshal(int32_t p0, nbyteslice p1, nint p2);

extern int32_t new_holochain_Header();

extern int32_t proxyholochain_Holochain_Activate(int32_t p0);

extern int32_t proxyholochain_Holochain_Agent(int32_t p0);

extern int32_t proxyholochain_Holochain_BSget(int32_t p0);

extern int32_t proxyholochain_Holochain_BSpost(int32_t p0);

extern int32_t proxyholochain_Holochain_BuildBridge(int32_t p0, int32_t p1, nstring p2);

extern int32_t proxyholochain_Holochain_Chain(int32_t p0);

extern void proxyholochain_Holochain_Close(int32_t p0);

extern nstring proxyholochain_Holochain_DBPath(int32_t p0);

extern int32_t proxyholochain_Holochain_DHT(int32_t p0);

extern nstring proxyholochain_Holochain_DNAPath(int32_t p0);

extern void proxyholochain_Holochain_Debug(int32_t p0, nstring p1);

/* Return type for proxyholochain_Holochain_GetEntryDef */
struct proxyholochain_Holochain_GetEntryDef_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Holochain_GetEntryDef_return proxyholochain_Holochain_GetEntryDef(int32_t p0, nstring p1);

/* Return type for proxyholochain_Holochain_GetProperty */
struct proxyholochain_Holochain_GetProperty_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_Holochain_GetProperty_return proxyholochain_Holochain_GetProperty(int32_t p0, nstring p1);

/* Return type for proxyholochain_Holochain_GetZome */
struct proxyholochain_Holochain_GetZome_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Holochain_GetZome_return proxyholochain_Holochain_GetZome(int32_t p0, nstring p1);

/* Return type for proxyholochain_Holochain_GetZomeForEntryType */
struct proxyholochain_Holochain_GetZomeForEntryType_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Holochain_GetZomeForEntryType_return proxyholochain_Holochain_GetZomeForEntryType(int32_t p0, nstring p1);

extern int32_t proxyholochain_Holochain_HandleAsyncSends(int32_t p0);

/* Return type for proxyholochain_Holochain_MakeRibosome */
struct proxyholochain_Holochain_MakeRibosome_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Holochain_MakeRibosome_return proxyholochain_Holochain_MakeRibosome(int32_t p0, nstring p1);

extern nstring proxyholochain_Holochain_Name(int32_t p0);

extern int32_t proxyholochain_Holochain_Node(int32_t p0);

extern nstring proxyholochain_Holochain_NodeIDStr(int32_t p0);

extern int32_t proxyholochain_Holochain_Nucleus(int32_t p0);

extern int32_t proxyholochain_Holochain_Prepare(int32_t p0);

extern int32_t proxyholochain_Holochain_PrepareHashType(int32_t p0);

extern int32_t proxyholochain_Holochain_Reset(int32_t p0);

extern nstring proxyholochain_Holochain_RootPath(int32_t p0);

/* Return type for proxyholochain_Holochain_Sign */
struct proxyholochain_Holochain_Sign_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_Holochain_Sign_return proxyholochain_Holochain_Sign(int32_t p0, nbyteslice p1);

extern void proxyholochain_Holochain_StartBackgroundTasks(int32_t p0);

extern char proxyholochain_Holochain_Started(int32_t p0);

extern nstring proxyholochain_Holochain_TestPath(int32_t p0);

extern nstring proxyholochain_Holochain_UIPath(int32_t p0);

extern nstring proxyholochain_Holochain_ZomePath(int32_t p0, int32_t p1);

extern int32_t new_holochain_Holochain();

/* Return type for proxyholochain_JSONEntry_Marshal */
struct proxyholochain_JSONEntry_Marshal_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_JSONEntry_Marshal_return proxyholochain_JSONEntry_Marshal(int32_t p0);

extern int32_t proxyholochain_JSONEntry_Unmarshal(int32_t p0, nbyteslice p1);

extern int32_t new_holochain_JSONEntry();

extern int32_t new_holochain_JSONSchemaValidator();

extern int32_t proxyholochain_JSRibosome_ChainGenesis(int32_t p0);

/* Return type for proxyholochain_JSRibosome_Receive */
struct proxyholochain_JSRibosome_Receive_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_JSRibosome_Receive_return proxyholochain_JSRibosome_Receive(int32_t p0, nstring p1, nstring p2);

extern nstring proxyholochain_JSRibosome_Type(int32_t p0);

extern int32_t new_holochain_JSRibosome();

extern int32_t new_holochain_LibP2PAgent();

extern void proxyholochain_Link_LinkAction_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Link_LinkAction_Get(int32_t p0);

extern void proxyholochain_Link_Base_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Link_Base_Get(int32_t p0);

extern void proxyholochain_Link_Link_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Link_Link_Get(int32_t p0);

extern void proxyholochain_Link_Tag_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Link_Tag_Get(int32_t p0);

extern int32_t new_holochain_Link();

extern void proxyholochain_LinkEvent_Status_Set(int32_t p0, nint p1);

extern nint proxyholochain_LinkEvent_Status_Get(int32_t p0);

extern void proxyholochain_LinkEvent_Source_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_LinkEvent_Source_Get(int32_t p0);

extern void proxyholochain_LinkEvent_LinksEntry_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_LinkEvent_LinksEntry_Get(int32_t p0);

extern int32_t new_holochain_LinkEvent();

extern void proxyholochain_LinkQuery_T_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_LinkQuery_T_Get(int32_t p0);

extern void proxyholochain_LinkQuery_StatusMask_Set(int32_t p0, nint p1);

extern nint proxyholochain_LinkQuery_StatusMask_Get(int32_t p0);

extern int32_t new_holochain_LinkQuery();

extern int32_t new_holochain_LinkQueryResp();

extern int32_t new_holochain_LinkReq();

extern int32_t new_holochain_LinksEntry();

extern void proxyholochain_ListAddReq_ListType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ListAddReq_ListType_Get(int32_t p0);

extern void proxyholochain_ListAddReq_WarrantType_Set(int32_t p0, nint p1);

extern nint proxyholochain_ListAddReq_WarrantType_Get(int32_t p0);

extern void proxyholochain_ListAddReq_Warrant_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_ListAddReq_Warrant_Get(int32_t p0);

extern int32_t new_holochain_ListAddReq();

extern void proxyholochain_Logger_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Logger_Name_Get(int32_t p0);

extern void proxyholochain_Logger_Enabled_Set(int32_t p0, char p1);

extern char proxyholochain_Logger_Enabled_Get(int32_t p0);

extern void proxyholochain_Logger_Format_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Logger_Format_Get(int32_t p0);

extern void proxyholochain_Logger_Prefix_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Logger_Prefix_Get(int32_t p0);

extern void proxyholochain_Logger_SetPrefix(int32_t p0, nstring p1);

extern int32_t new_holochain_Logger();

extern int32_t new_holochain_Loggers();

/* Return type for proxyholochain_Message_Encode */
struct proxyholochain_Message_Encode_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_Message_Encode_return proxyholochain_Message_Encode(int32_t p0);

extern nstring proxyholochain_Message_String(int32_t p0);

extern int32_t new_holochain_Message();

extern void proxyholochain_Meta_T_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Meta_T_Get(int32_t p0);

extern void proxyholochain_Meta_V_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_Meta_V_Get(int32_t p0);

extern int32_t new_holochain_Meta();

extern void proxyholochain_ModAgentOptions_Identity_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ModAgentOptions_Identity_Get(int32_t p0);

extern void proxyholochain_ModAgentOptions_Revocation_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ModAgentOptions_Revocation_Get(int32_t p0);

extern int32_t new_holochain_ModAgentOptions();

extern int32_t new_holochain_ModReq();

extern int32_t proxyholochain_Node_Close(int32_t p0);

extern int32_t proxyholochain_Node_StartProtocol(int32_t p0, int32_t p1, nint p2);

extern int32_t new_holochain_Node();

extern int32_t proxyholochain_Nucleus_DNA(int32_t p0);

extern int32_t proxyholochain_Nucleus_RunGenesis(int32_t p0);

extern int32_t proxyholochain_Nucleus_Start(int32_t p0);

extern int32_t new_holochain_Nucleus();

extern void proxyholochain_Package_Chain_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_Package_Chain_Get(int32_t p0);

extern int32_t new_holochain_Package();

extern void proxyholochain_PeerInfo_ID_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_PeerInfo_ID_Get(int32_t p0);

extern int32_t new_holochain_PeerInfo();

extern int32_t new_holochain_PeerList();

extern void proxyholochain_PeerRecord_Warrant_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_PeerRecord_Warrant_Get(int32_t p0);

extern int32_t new_holochain_PeerRecord();

extern void proxyholochain_Progenitor_Identity_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Progenitor_Identity_Get(int32_t p0);

extern void proxyholochain_Progenitor_PubKey_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_Progenitor_PubKey_Get(int32_t p0);

extern int32_t new_holochain_Progenitor();

extern int32_t new_holochain_Protocol();

extern void proxyholochain_Put_Idx_Set(int32_t p0, nint p1);

extern nint proxyholochain_Put_Idx_Get(int32_t p0);

extern int32_t new_holochain_Put();

extern void proxyholochain_PutReq_S_Set(int32_t p0, nint p1);

extern nint proxyholochain_PutReq_S_Get(int32_t p0);

extern int32_t new_holochain_PutReq();

extern void proxyholochain_QueryConstrain_Contains_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_QueryConstrain_Contains_Get(int32_t p0);

extern void proxyholochain_QueryConstrain_Equals_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_QueryConstrain_Equals_Get(int32_t p0);

extern void proxyholochain_QueryConstrain_Matches_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_QueryConstrain_Matches_Get(int32_t p0);

extern void proxyholochain_QueryConstrain_Count_Set(int32_t p0, nint p1);

extern nint proxyholochain_QueryConstrain_Count_Get(int32_t p0);

extern void proxyholochain_QueryConstrain_Page_Set(int32_t p0, nint p1);

extern nint proxyholochain_QueryConstrain_Page_Get(int32_t p0);

extern int32_t new_holochain_QueryConstrain();

extern int32_t new_holochain_QueryOptions();

extern void proxyholochain_QueryOrder_Ascending_Set(int32_t p0, char p1);

extern char proxyholochain_QueryOrder_Ascending_Get(int32_t p0);

extern int32_t new_holochain_QueryOrder();

extern void proxyholochain_QueryResult_Header_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_QueryResult_Header_Get(int32_t p0);

extern void proxyholochain_QueryResult_Entry_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_QueryResult_Entry_Get(int32_t p0);

extern int32_t new_holochain_QueryResult();

extern void proxyholochain_QueryReturn_Hashes_Set(int32_t p0, char p1);

extern char proxyholochain_QueryReturn_Hashes_Get(int32_t p0);

extern void proxyholochain_QueryReturn_Entries_Set(int32_t p0, char p1);

extern char proxyholochain_QueryReturn_Entries_Get(int32_t p0);

extern void proxyholochain_QueryReturn_Headers_Set(int32_t p0, char p1);

extern char proxyholochain_QueryReturn_Headers_Get(int32_t p0);

extern int32_t new_holochain_QueryReturn();

extern char proxyholochain_RoutingTable_IsEmpty(int32_t p0);

extern void proxyholochain_RoutingTable_Print(int32_t p0);

extern nint proxyholochain_RoutingTable_Size(int32_t p0);

extern int32_t new_holochain_RoutingTable();

extern void proxyholochain_SelfRevocation_Data_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_SelfRevocation_Data_Get(int32_t p0);

extern void proxyholochain_SelfRevocation_OldSig_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_SelfRevocation_OldSig_Get(int32_t p0);

extern void proxyholochain_SelfRevocation_NewSig_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_SelfRevocation_NewSig_Get(int32_t p0);

/* Return type for proxyholochain_SelfRevocation_Marshal */
struct proxyholochain_SelfRevocation_Marshal_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_SelfRevocation_Marshal_return proxyholochain_SelfRevocation_Marshal(int32_t p0);

extern int32_t proxyholochain_SelfRevocation_Unmarshal(int32_t p0, nbyteslice p1);

extern int32_t proxyholochain_SelfRevocation_Verify(int32_t p0);

extern int32_t new_holochain_SelfRevocation();

extern int32_t proxyholochain_SelfRevocationWarrant_Decode(int32_t p0, nbyteslice p1);

/* Return type for proxyholochain_SelfRevocationWarrant_Encode */
struct proxyholochain_SelfRevocationWarrant_Encode_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_SelfRevocationWarrant_Encode_return proxyholochain_SelfRevocationWarrant_Encode(int32_t p0);

extern nint proxyholochain_SelfRevocationWarrant_Type(int32_t p0);

extern int32_t proxyholochain_SelfRevocationWarrant_Verify(int32_t p0, int32_t p1);

extern int32_t new_holochain_SelfRevocationWarrant();

extern void proxyholochain_SendOptions_Callback_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_SendOptions_Callback_Get(int32_t p0);

extern void proxyholochain_SendOptions_Timeout_Set(int32_t p0, nint p1);

extern nint proxyholochain_SendOptions_Timeout_Get(int32_t p0);

extern int32_t new_holochain_SendOptions();

extern void proxyholochain_Service_DefaultAgent_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_Service_DefaultAgent_Get(int32_t p0);

extern void proxyholochain_Service_Path_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Service_Path_Get(int32_t p0);

/* Return type for proxyholochain_Service_Clone */
struct proxyholochain_Service_Clone_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Service_Clone_return proxyholochain_Service_Clone(int32_t p0, nstring p1, nstring p2, int32_t p3, char p4, char p5);

/* Return type for proxyholochain_Service_GenChain */
struct proxyholochain_Service_GenChain_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Service_GenChain_return proxyholochain_Service_GenChain(int32_t p0, nstring p1);

extern int32_t proxyholochain_Service_InitAppDir(int32_t p0, nstring p1, nstring p2);

/* Return type for proxyholochain_Service_IsConfigured */
struct proxyholochain_Service_IsConfigured_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_Service_IsConfigured_return proxyholochain_Service_IsConfigured(int32_t p0, nstring p1);

extern nstring proxyholochain_Service_ListChains(int32_t p0);

/* Return type for proxyholochain_Service_Load */
struct proxyholochain_Service_Load_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Service_Load_return proxyholochain_Service_Load(int32_t p0, nstring p1);

/* Return type for proxyholochain_Service_MakeAppPackage */
struct proxyholochain_Service_MakeAppPackage_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_Service_MakeAppPackage_return proxyholochain_Service_MakeAppPackage(int32_t p0, int32_t p1);

/* Return type for proxyholochain_Service_MakeTestingApp */
struct proxyholochain_Service_MakeTestingApp_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Service_MakeTestingApp_return proxyholochain_Service_MakeTestingApp(int32_t p0, nstring p1, nstring p2, char p3, char p4, int32_t p5);

extern int32_t new_holochain_Service();

extern void proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set(int32_t p0, char p1);

extern char proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get(int32_t p0);

extern void proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set(int32_t p0, char p1);

extern char proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get(int32_t p0);

extern void proxyholochain_ServiceConfig_DefaultBootstrapServer_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ServiceConfig_DefaultBootstrapServer_Get(int32_t p0);

extern void proxyholochain_ServiceConfig_DefaultEnableMDNS_Set(int32_t p0, char p1);

extern char proxyholochain_ServiceConfig_DefaultEnableMDNS_Get(int32_t p0);

extern void proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set(int32_t p0, char p1);

extern char proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get(int32_t p0);

extern int32_t proxyholochain_ServiceConfig_Validate(int32_t p0);

extern int32_t new_holochain_ServiceConfig();

extern void proxyholochain_Signature_S_Set(int32_t p0, nbyteslice p1);

extern nbyteslice proxyholochain_Signature_S_Get(int32_t p0);

extern int32_t new_holochain_Signature();

extern void proxyholochain_StatusChange_Action_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_StatusChange_Action_Get(int32_t p0);

extern int32_t new_holochain_StatusChange();

extern void proxyholochain_TaggedHash_H_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TaggedHash_H_Get(int32_t p0);

extern void proxyholochain_TaggedHash_E_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TaggedHash_E_Get(int32_t p0);

extern void proxyholochain_TaggedHash_EntryType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TaggedHash_EntryType_Get(int32_t p0);

extern void proxyholochain_TaggedHash_T_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TaggedHash_T_Get(int32_t p0);

extern void proxyholochain_TaggedHash_Source_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TaggedHash_Source_Get(int32_t p0);

extern int32_t new_holochain_TaggedHash();

extern void proxyholochain_TestConfig_GossipInterval_Set(int32_t p0, nint p1);

extern nint proxyholochain_TestConfig_GossipInterval_Get(int32_t p0);

extern void proxyholochain_TestConfig_Duration_Set(int32_t p0, nint p1);

extern nint proxyholochain_TestConfig_Duration_Get(int32_t p0);

extern int32_t new_holochain_TestConfig();

extern void proxyholochain_TestData_Convey_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TestData_Convey_Get(int32_t p0);

extern void proxyholochain_TestData_Zome_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TestData_Zome_Get(int32_t p0);

extern void proxyholochain_TestData_FnName_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TestData_FnName_Get(int32_t p0);

extern void proxyholochain_TestData_Regexp_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TestData_Regexp_Get(int32_t p0);

extern void proxyholochain_TestData_Exposure_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TestData_Exposure_Get(int32_t p0);

extern void proxyholochain_TestData_Raw_Set(int32_t p0, char p1);

extern char proxyholochain_TestData_Raw_Get(int32_t p0);

extern void proxyholochain_TestData_Repeat_Set(int32_t p0, nint p1);

extern nint proxyholochain_TestData_Repeat_Get(int32_t p0);

extern void proxyholochain_TestData_Benchmark_Set(int32_t p0, char p1);

extern char proxyholochain_TestData_Benchmark_Get(int32_t p0);

extern int32_t new_holochain_TestData();

extern int32_t new_holochain_TestFixtures();

extern void proxyholochain_TestSet_Identity_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_TestSet_Identity_Get(int32_t p0);

extern void proxyholochain_TestSet_Benchmark_Set(int32_t p0, char p1);

extern char proxyholochain_TestSet_Benchmark_Get(int32_t p0);

extern int32_t new_holochain_TestSet();

extern int32_t new_holochain_ValidateQuery();

extern void proxyholochain_ValidateResponse_Type_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ValidateResponse_Type_Get(int32_t p0);

extern int32_t new_holochain_ValidateResponse();

extern void proxyholochain_ValidationPackage_Chain_Set(int32_t p0, int32_t p1);

extern int32_t proxyholochain_ValidationPackage_Chain_Get(int32_t p0);

extern int32_t new_holochain_ValidationPackage();

extern void proxyholochain_Zome_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Zome_Name_Get(int32_t p0);

extern void proxyholochain_Zome_Description_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Zome_Description_Get(int32_t p0);

extern void proxyholochain_Zome_Code_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Zome_Code_Get(int32_t p0);

extern void proxyholochain_Zome_RibosomeType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_Zome_RibosomeType_Get(int32_t p0);

extern nstring proxyholochain_Zome_CodeFileName(int32_t p0);

/* Return type for proxyholochain_Zome_GetEntryDef */
struct proxyholochain_Zome_GetEntryDef_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Zome_GetEntryDef_return proxyholochain_Zome_GetEntryDef(int32_t p0, nstring p1);

/* Return type for proxyholochain_Zome_GetFunctionDef */
struct proxyholochain_Zome_GetFunctionDef_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Zome_GetFunctionDef_return proxyholochain_Zome_GetFunctionDef(int32_t p0, nstring p1);

/* Return type for proxyholochain_Zome_MakeRibosome */
struct proxyholochain_Zome_MakeRibosome_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain_Zome_MakeRibosome_return proxyholochain_Zome_MakeRibosome(int32_t p0, int32_t p1);

extern int32_t new_holochain_Zome();

extern void proxyholochain_ZomeFile_Name_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ZomeFile_Name_Get(int32_t p0);

extern void proxyholochain_ZomeFile_Description_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ZomeFile_Description_Get(int32_t p0);

extern void proxyholochain_ZomeFile_CodeFile_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ZomeFile_CodeFile_Get(int32_t p0);

extern void proxyholochain_ZomeFile_RibosomeType_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ZomeFile_RibosomeType_Get(int32_t p0);

extern void proxyholochain_ZomeFile_BridgeTo_Set(int32_t p0, nstring p1);

extern nstring proxyholochain_ZomeFile_BridgeTo_Get(int32_t p0);

extern int32_t new_holochain_ZomeFile();

extern int32_t proxyholochain_ZygoRibosome_ChainGenesis(int32_t p0);

/* Return type for proxyholochain_ZygoRibosome_Receive */
struct proxyholochain_ZygoRibosome_Receive_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_ZygoRibosome_Receive_return proxyholochain_ZygoRibosome_Receive(int32_t p0, nstring p1, nstring p2);

extern nstring proxyholochain_ZygoRibosome_Type(int32_t p0);

extern int32_t new_holochain_ZygoRibosome();

extern nstring proxyholochain_Action_Name(int32_t p0);

extern int32_t proxyholochain_CommittingAction_CheckValidationRequest(int32_t p0, int32_t p1);

extern int32_t proxyholochain_CommittingAction_Entry(int32_t p0);

extern nstring proxyholochain_CommittingAction_EntryType(int32_t p0);

extern nstring proxyholochain_CommittingAction_Name(int32_t p0);

extern void proxyholochain_CommittingAction_SetHeader(int32_t p0, int32_t p1);

/* Return type for proxyholochain_Entry_Marshal */
struct proxyholochain_Entry_Marshal_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_Entry_Marshal_return proxyholochain_Entry_Marshal(int32_t p0);

extern int32_t proxyholochain_Entry_Unmarshal(int32_t p0, nbyteslice p1);

/* Return type for proxyholochain_Revocation_Marshal */
struct proxyholochain_Revocation_Marshal_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_Revocation_Marshal_return proxyholochain_Revocation_Marshal(int32_t p0);

extern int32_t proxyholochain_Revocation_Unmarshal(int32_t p0, nbyteslice p1);

extern int32_t proxyholochain_Revocation_Verify(int32_t p0);

extern int32_t proxyholochain_Ribosome_ChainGenesis(int32_t p0);

/* Return type for proxyholochain_Ribosome_Receive */
struct proxyholochain_Ribosome_Receive_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain_Ribosome_Receive_return proxyholochain_Ribosome_Receive(int32_t p0, nstring p1, nstring p2);

extern nstring proxyholochain_Ribosome_Type(int32_t p0);

extern int32_t proxyholochain_ValidatingAction_CheckValidationRequest(int32_t p0, int32_t p1);

extern nstring proxyholochain_ValidatingAction_Name(int32_t p0);

// skipped method ValidatingAction.Receive with unsupported parameter or result types
// skipped method ValidatingAction.SysValidation with unsupported parameter or result types

extern int32_t proxyholochain_Warrant_Decode(int32_t p0, nbyteslice p1);

/* Return type for proxyholochain_Warrant_Encode */
struct proxyholochain_Warrant_Encode_return {
	nbyteslice r0;
	int32_t r1;
};

extern struct proxyholochain_Warrant_Encode_return proxyholochain_Warrant_Encode(int32_t p0);

extern nint proxyholochain_Warrant_Type(int32_t p0);

extern int32_t proxyholochain_Warrant_Verify(int32_t p0, int32_t p1);

extern void var_setholochain_AgentEntryDef(int32_t p0);

extern int32_t var_getholochain_AgentEntryDef();

extern void var_setholochain_AlphaValue(nint p0);

extern nint var_getholochain_AlphaValue();

extern void var_setholochain_BasicTemplateAppPackage(nstring p0);

extern nstring var_getholochain_BasicTemplateAppPackage();

extern void var_setholochain_BridgeAppNotFoundErr(int32_t p0);

extern int32_t var_getholochain_BridgeAppNotFoundErr();

extern void var_setholochain_CapabilityInvalidErr(int32_t p0);

extern int32_t var_getholochain_CapabilityInvalidErr();

extern void var_setholochain_CloserPeerCount(nint p0);

extern nint var_getholochain_CloserPeerCount();

extern void var_setholochain_DNAEntryDef(int32_t p0);

extern int32_t var_getholochain_DNAEntryDef();

extern void var_setholochain_EnableAllLoggersEnv(nstring p0);

extern nstring var_getholochain_EnableAllLoggersEnv();

extern void var_setholochain_ErrBlockedListed(int32_t p0);

extern int32_t var_getholochain_ErrBlockedListed();

extern void var_setholochain_ErrDHTErrNoGossipersAvailable(int32_t p0);

extern int32_t var_getholochain_ErrDHTErrNoGossipersAvailable();

extern void var_setholochain_ErrDHTExpectedGossipReqInBody(int32_t p0);

extern int32_t var_getholochain_ErrDHTExpectedGossipReqInBody();

extern void var_setholochain_ErrDHTUnexpectedTypeInBody(int32_t p0);

extern int32_t var_getholochain_ErrDHTUnexpectedTypeInBody();

extern void var_setholochain_ErrEmptyRoutingTable(int32_t p0);

extern int32_t var_getholochain_ErrEmptyRoutingTable();

extern void var_setholochain_ErrEntryTypeMismatch(int32_t p0);

extern int32_t var_getholochain_ErrEntryTypeMismatch();

extern void var_setholochain_ErrHashDeleted(int32_t p0);

extern int32_t var_getholochain_ErrHashDeleted();

extern void var_setholochain_ErrHashModified(int32_t p0);

extern int32_t var_getholochain_ErrHashModified();

extern void var_setholochain_ErrHashNotFound(int32_t p0);

extern int32_t var_getholochain_ErrHashNotFound();

extern void var_setholochain_ErrHashRejected(int32_t p0);

extern int32_t var_getholochain_ErrHashRejected();

extern void var_setholochain_ErrIncompleteChain(int32_t p0);

extern int32_t var_getholochain_ErrIncompleteChain();

extern void var_setholochain_ErrLinkNotFound(int32_t p0);

extern int32_t var_getholochain_ErrLinkNotFound();

extern void var_setholochain_ErrNilEntryInvalid(int32_t p0);

extern int32_t var_getholochain_ErrNilEntryInvalid();

extern void var_setholochain_ErrNoSuchIdx(int32_t p0);

extern int32_t var_getholochain_ErrNoSuchIdx();

extern void var_setholochain_ErrNotValidForAgentType(int32_t p0);

extern int32_t var_getholochain_ErrNotValidForAgentType();

extern void var_setholochain_ErrNotValidForDNAType(int32_t p0);

extern int32_t var_getholochain_ErrNotValidForDNAType();

extern void var_setholochain_ErrNotValidForKeyType(int32_t p0);

extern int32_t var_getholochain_ErrNotValidForKeyType();

extern void var_setholochain_ErrPutLinkOverDeleted(int32_t p0);

extern int32_t var_getholochain_ErrPutLinkOverDeleted();

extern void var_setholochain_ErrWrongNargs(int32_t p0);

extern int32_t var_getholochain_ErrWrongNargs();

extern void var_setholochain_IsDevMode(char p0);

extern char var_getholochain_IsDevMode();

extern void var_setholochain_KValue(nint p0);

extern nint var_getholochain_KValue();

extern void var_setholochain_KeyEntryDef(int32_t p0);

extern int32_t var_getholochain_KeyEntryDef();

extern void var_setholochain_NonCallableAction(int32_t p0);

extern int32_t var_getholochain_NonCallableAction();

extern void var_setholochain_NonDHTAction(int32_t p0);

extern int32_t var_getholochain_NonDHTAction();

extern void var_setholochain_SelfRevocationDoesNotVerify(int32_t p0);

extern int32_t var_getholochain_SelfRevocationDoesNotVerify();

extern void var_setholochain_SendTimeoutErr(int32_t p0);

extern int32_t var_getholochain_SendTimeoutErr();

extern void var_setholochain_UnknownWarrantTypeErr(int32_t p0);

extern int32_t var_getholochain_UnknownWarrantTypeErr();

extern void var_setholochain_ValidationFailedErr(int32_t p0);

extern int32_t var_getholochain_ValidationFailedErr();

extern void var_setholochain_WarrantPropertyNotFoundErr(int32_t p0);

extern int32_t var_getholochain_WarrantPropertyNotFoundErr();

// skipped function ActionReceiver with unsupported parameter or result types

extern void proxyholochain__BootstrapRefreshTask(int32_t p0);

/* Return type for proxyholochain__BuildJSONSchemaValidatorFromFile */
struct proxyholochain__BuildJSONSchemaValidatorFromFile_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain__BuildJSONSchemaValidatorFromFile_return proxyholochain__BuildJSONSchemaValidatorFromFile(nstring p0);

/* Return type for proxyholochain__BuildJSONSchemaValidatorFromString */
struct proxyholochain__BuildJSONSchemaValidatorFromString_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain__BuildJSONSchemaValidatorFromString_return proxyholochain__BuildJSONSchemaValidatorFromString(nstring p0);

// skipped function ByteDecoder with unsupported parameter or result types
// skipped function ByteEncoder with unsupported parameter or result types

extern int32_t proxyholochain__CopyDir(nstring p0, nstring p1);

extern int32_t proxyholochain__CopyFile(nstring p0, nstring p1);

/* Return type for proxyholochain__CreateRibosome */
struct proxyholochain__CreateRibosome_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain__CreateRibosome_return proxyholochain__CreateRibosome(int32_t p0, int32_t p1);

// skipped function DNAHashofUngenedChain with unsupported parameter or result types

extern void proxyholochain__Debug(nstring p0);

/* Return type for proxyholochain__DecodeWarrant */
struct proxyholochain__DecodeWarrant_return {
	int32_t r0;
	int32_t r1;
};

// skipped function Debugf with unsupported parameter or result types
// skipped function Decode with unsupported parameter or result types

extern struct proxyholochain__DecodeWarrant_return proxyholochain__DecodeWarrant(nint p0, nbyteslice p1);

// skipped function DirExists with unsupported parameter or result types
// skipped function Encode with unsupported parameter or result types

extern nstring proxyholochain__EncodingFormat(nstring p0);

extern nstring proxyholochain__EscapeJSONValue(nstring p0);

// skipped function FileExists with unsupported parameter or result types
// skipped function FileSize with unsupported parameter or result types
// skipped function GetAllTestRoles with unsupported parameter or result types
// skipped function GetTestScenarioRoles with unsupported parameter or result types
// skipped function GetTestScenarios with unsupported parameter or result types
// skipped function GossipReceiver with unsupported parameter or result types

extern void proxyholochain__GossipTask(int32_t p0);

extern void proxyholochain__Info(nstring p0);

// skipped function Infof with unsupported parameter or result types
// skipped function Init with unsupported parameter or result types

extern void proxyholochain__InitializeHolochain();

extern char proxyholochain__IsInitialized(nstring p0);

/* Return type for proxyholochain__LoadAgent */
struct proxyholochain__LoadAgent_return {
	int32_t r0;
	int32_t r1;
};

// skipped function KademliaReceiver with unsupported parameter or result types

extern struct proxyholochain__LoadAgent_return proxyholochain__LoadAgent(nstring p0);

/* Return type for proxyholochain__LoadService */
struct proxyholochain__LoadService_return {
	int32_t r0;
	int32_t r1;
};

// skipped function LoadAppPackage with unsupported parameter or result types

extern struct proxyholochain__LoadService_return proxyholochain__LoadService(nstring p0);

/* Return type for proxyholochain__LoadTestConfig */
struct proxyholochain__LoadTestConfig_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain__LoadTestConfig_return proxyholochain__LoadTestConfig(nstring p0);

/* Return type for proxyholochain__MakeActionFromMessage */
struct proxyholochain__MakeActionFromMessage_return {
	int32_t r0;
	int32_t r1;
};

// skipped function LoadTestFile with unsupported parameter or result types
// skipped function LoadTestFiles with unsupported parameter or result types

extern struct proxyholochain__MakeActionFromMessage_return proxyholochain__MakeActionFromMessage(int32_t p0);

extern int32_t proxyholochain__MakeDirs(nstring p0);

/* Return type for proxyholochain__MakeValidationPackage */
struct proxyholochain__MakeValidationPackage_return {
	int32_t r0;
	int32_t r1;
};

// skipped function MakePackage with unsupported parameter or result types

extern struct proxyholochain__MakeValidationPackage_return proxyholochain__MakeValidationPackage(int32_t p0, int32_t p1);

// skipped function MarshalEntry with unsupported parameter or result types
// skipped function MarshalHeader with unsupported parameter or result types
// skipped function MarshalSignature with unsupported parameter or result types
// skipped function NewAgent with unsupported parameter or result types
// skipped function NewBridgeAction with unsupported parameter or result types
// skipped function NewCallAction with unsupported parameter or result types
// skipped function NewCapability with unsupported parameter or result types
// skipped function NewChain with unsupported parameter or result types
// skipped function NewChainFromFile with unsupported parameter or result types

extern int32_t proxyholochain__NewCommitAction(nstring p0, int32_t p1);

extern int32_t proxyholochain__NewDHT(int32_t p0);

extern int32_t proxyholochain__NewDebugAction(nstring p0);

// skipped function NewDelAction with unsupported parameter or result types
// skipped function NewErrorResponse with unsupported parameter or result types
// skipped function NewGetAction with unsupported parameter or result types

extern int32_t proxyholochain__NewGetBridgesAction(nbyteslice p0);

extern int32_t proxyholochain__NewGetLinksAction(int32_t p0, int32_t p1);

/* Return type for proxyholochain__NewJSRibosome */
struct proxyholochain__NewJSRibosome_return {
	int32_t r0;
	int32_t r1;
};

// skipped function NewHolochain with unsupported parameter or result types

extern struct proxyholochain__NewJSRibosome_return proxyholochain__NewJSRibosome(int32_t p0, int32_t p1);

// skipped function NewLinkAction with unsupported parameter or result types
// skipped function NewListAddAction with unsupported parameter or result types

extern int32_t proxyholochain__NewMakeHashAction(int32_t p0);

/* Return type for proxyholochain__NewNode */
struct proxyholochain__NewNode_return {
	int32_t r0;
	int32_t r1;
};

// skipped function NewModAction with unsupported parameter or result types
// skipped function NewModAgentAction with unsupported parameter or result types

extern struct proxyholochain__NewNode_return proxyholochain__NewNode(nstring p0, nstring p1, int32_t p2, char p3, int32_t p4);

extern int32_t proxyholochain__NewNucleus(int32_t p0, int32_t p1);

extern int32_t proxyholochain__NewPropertyAction(nstring p0);

extern int32_t proxyholochain__NewPutAction(nstring p0, int32_t p1, int32_t p2);

extern int32_t proxyholochain__NewQueryAction(int32_t p0);

/* Return type for proxyholochain__NewSelfRevocationWarrant */
struct proxyholochain__NewSelfRevocationWarrant_return {
	int32_t r0;
	int32_t r1;
};

// skipped function NewRoutingTable with unsupported parameter or result types
// skipped function NewSelfRevocation with unsupported parameter or result types

extern struct proxyholochain__NewSelfRevocationWarrant_return proxyholochain__NewSelfRevocationWarrant(int32_t p0);

// skipped function NewSendAction with unsupported parameter or result types

extern int32_t proxyholochain__NewSignAction(nbyteslice p0);

extern int32_t proxyholochain__NewVerifySignatureAction(nstring p0, nstring p1, nstring p2);

/* Return type for proxyholochain__NewZygoRibosome */
struct proxyholochain__NewZygoRibosome_return {
	int32_t r0;
	int32_t r1;
};

extern struct proxyholochain__NewZygoRibosome_return proxyholochain__NewZygoRibosome(int32_t p0, int32_t p1);

/* Return type for proxyholochain__PrettyPrintJSON */
struct proxyholochain__PrettyPrintJSON_return {
	nstring r0;
	int32_t r1;
};

extern struct proxyholochain__PrettyPrintJSON_return proxyholochain__PrettyPrintJSON(nbyteslice p0);

// skipped function ReadFile with unsupported parameter or result types

extern void proxyholochain__RegisterBultinRibosomes();

// skipped function RegisterRibosome with unsupported parameter or result types

extern void proxyholochain__RetryTask(int32_t p0);

extern void proxyholochain__RoutingRefreshTask(int32_t p0);

// skipped function RunValidationPhase with unsupported parameter or result types

extern int32_t proxyholochain__SaveAgent(nstring p0, int32_t p1);

// skipped function SortClosestPeers with unsupported parameter or result types

extern nstring proxyholochain__TestingAppAppPackage();

extern nstring proxy_error_Error(int32_t p0);

// IncGoRef is called by foreign code to pin a Go object while its refnum is crossing
// the language barrier

extern void IncGoRef(int32_t p0);

// DestroyRef is called by Java to inform Go it is done with a reference.

extern void DestroyRef(int32_t p0);

#ifdef __cplusplus
}
#endif
