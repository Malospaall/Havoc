package agent

const (
	DEMON_MAGIC_VALUE = 0xDEADBEEF
)

const (
	/*
	 * https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
	 * Pipe write operations across a network are limited in size per write.
	 * The amount varies per platform. For x86 platforms it's 63.97 MB.
	 * For x64 platforms it's 31.97 MB. For Itanium it's 63.95 MB.
	 */
	// we are using 30 MB
	DEMON_MAX_RESPONSE_LENGTH = 0x1e00000
)

const (
	PROCESS_ARCH_UNKNOWN = 0
	PROCESS_ARCH_X86     = 1
	PROCESS_ARCH_X64     = 2
	PROCESS_ARCH_IA64    = 3
)

// TODO: change Command IDs. use something more readable and understandable.
const (
	COMMAND_GET_JOB                 = 1
	DEMON_INIT                      = 99
	COMMAND_CHECKIN                 = 100
	COMMAND_NOJOB                   = 10
	COMMAND_SLEEP                   = 11
	COMMAND_PROC                    = 0x1010
	COMMAND_PS_IMPORT               = 0x1011
	COMMAND_PROC_LIST               = 12
	COMMAND_FS                      = 15
	COMMAND_INLINEEXECUTE           = 20
	COMMAND_ASSEMBLY_INLINE_EXECUTE = 0x2001
	COMMAND_ASSEMBLY_LIST_VERSIONS  = 0x2003
	COMMAND_JOB                     = 21
	COMMAND_INJECT_DLL              = 22
	COMMAND_INJECT_SHELLCODE        = 24
	COMMAND_SPAWNDLL                = 26
	COMMAND_PROC_PPIDSPOOF          = 27
	COMMAND_TOKEN                   = 40
	COMMAND_NET                     = 2100
	COMMAND_CONFIG                  = 2500
	COMMAND_SCREENSHOT              = 2510
	COMMAND_PIVOT                   = 2520
	COMMAND_TRANSFER                = 2530
	COMMAND_SOCKET                  = 2540
	COMMAND_KERBEROS                = 2550
	COMMAND_MEM_FILE                = 2560
	COMMAND_PACKAGE_DROPPED         = 2570

	DEMON_INFO = 89

	COMMAND_OUTPUT    = 90
	COMMAND_OUTPUTW   = 9000
	COMMAND_ERROR     = 91
	COMMAND_EXIT      = 92
	COMMAND_KILL_DATE = 93
	BEACON_OUTPUT     = 94

	COMMAND_INLINEEXECUTE_EXCEPTION        = 1
	COMMAND_INLINEEXECUTE_SYMBOL_NOT_FOUND = 2
	COMMAND_INLINEEXECUTE_RAN_OK           = 3
	COMMAND_INLINEEXECUTE_COULD_NO_RUN     = 4

	COMMAND_EXCEPTION        = 0x98
	COMMAND_SYMBOL_NOT_FOUND = 0x99

	CALLBACK_OUTPUT      = 0x0
	CALLBACK_OUTPUT_OEM  = 0x1e
	CALLBACK_ERROR       = 0x0d
	CALLBACK_OUTPUT_UTF8 = 0x20
	CALLBACK_FILE        = 0x02
	CALLBACK_FILE_WRITE  = 0x08
	CALLBACK_FILE_CLOSE  = 0x09
)

const (
	CONFIG_IMPLANT_SPFTHREADSTART  = 3
	CONFIG_IMPLANT_SLEEP_TECHNIQUE = 5

	CONFIG_IMPLANT_VERBOSE         = 4
	CONFIG_IMPLANT_COFFEE_THREADED = 6
	CONFIG_IMPLANT_COFFEE_VEH      = 7

	CONFIG_MEMORY_ALLOC   = 101
	CONFIG_MEMORY_EXECUTE = 102

	CONFIG_INJECT_TECHNIQUE = 150
	CONFIG_INJECT_SPOOFADDR = 151
	CONFIG_INJECT_SPAWN64   = 152
	CONFIG_INJECT_SPAWN32   = 153

	CONFIG_KILLDATE     = 154
	CONFIG_WORKINGHOURS = 155

	DEMON_NET_COMMAND_DOMAIN     = 1
	DEMON_NET_COMMAND_LOGONS     = 2
	DEMON_NET_COMMAND_SESSIONS   = 3
	DEMON_NET_COMMAND_COMPUTER   = 4
	DEMON_NET_COMMAND_DCLIST     = 5
	DEMON_NET_COMMAND_SHARE      = 6
	DEMON_NET_COMMAND_LOCALGROUP = 7
	DEMON_NET_COMMAND_GROUP      = 8
	DEMON_NET_COMMAND_USERS      = 9

	DEMON_PIVOT_LIST           = 1
	DEMON_PIVOT_SMB_CONNECT    = 10
	DEMON_PIVOT_SMB_DISCONNECT = 11
	DEMON_PIVOT_SMB_COMMAND    = 12

	DEMON_INFO_MEM_ALLOC   = 10
	DEMON_INFO_MEM_EXEC    = 11
	DEMON_INFO_MEM_PROTECT = 12
	DEMON_INFO_PROC_CREATE = 21

	DEMON_COMMAND_JOB_LIST        = 1
	DEMON_COMMAND_JOB_SUSPEND     = 2
	DEMON_COMMAND_JOB_RESUME      = 3
	DEMON_COMMAND_JOB_KILL_REMOVE = 4
	DEMON_COMMAND_JOB_DIED        = 5

	DEMON_COMMAND_TRANSFER_LIST   = 0
	DEMON_COMMAND_TRANSFER_STOP   = 1
	DEMON_COMMAND_TRANSFER_RESUME = 2
	DEMON_COMMAND_TRANSFER_REMOVE = 3

	DEMON_COMMAND_PROC_MODULES = 2
	DEMON_COMMAND_PROC_GREP    = 3
	DEMON_COMMAND_PROC_CREATE  = 4
	DEMON_COMMAND_PROC_MEMORY  = 6
	DEMON_COMMAND_PROC_KILL    = 7

	DEMON_COMMAND_TOKEN_IMPERSONATE      = 1
	DEMON_COMMAND_TOKEN_STEAL            = 2
	DEMON_COMMAND_TOKEN_LIST             = 3
	DEMON_COMMAND_TOKEN_PRIVSGET_OR_LIST = 4
	DEMON_COMMAND_TOKEN_MAKE             = 5
	DEMON_COMMAND_TOKEN_GET_UID          = 6
	DEMON_COMMAND_TOKEN_REVERT           = 7
	DEMON_COMMAND_TOKEN_REMOVE           = 8
	DEMON_COMMAND_TOKEN_CLEAR            = 9
	DEMON_COMMAND_TOKEN_FIND_TOKENS      = 10

	DEMON_COMMAND_FS_DIR      = 1
	DEMON_COMMAND_FS_DOWNLOAD = 2
	DEMON_COMMAND_FS_UPLOAD   = 3
	DEMON_COMMAND_FS_CD       = 4
	DEMON_COMMAND_FS_REMOVE   = 5
	DEMON_COMMAND_FS_MKDIR    = 6
	DEMON_COMMAND_FS_COPY     = 7
        DEMON_COMMAND_FS_MOVE     = 8
	DEMON_COMMAND_FS_GET_PWD  = 9
	DEMON_COMMAND_FS_CAT      = 10
)

const (
	DOTNET_INFO_PATCHED     = 0x1
	DOTNET_INFO_NET_VERSION = 0x2
	DOTNET_INFO_ENTRYPOINT  = 0x3
	DOTNET_INFO_FINISHED    = 0x4
	DOTNET_INFO_FAILED      = 0x5
)

const (
	HAVOC_CONSOLE_MESSAGE = 0x80
	HAVOC_BOF_CALLBACK    = 0x81
)

const (
	ERROR_WIN32_LASTERROR = 1
	ERROR_TOKEN           = 3
)

const (
	SOCKET_COMMAND_RPORTFWD_ADD    = 0x0
	SOCKET_COMMAND_RPORTFWD_ADDLCL = 0x1
	SOCKET_COMMAND_RPORTFWD_LIST   = 0x2
	SOCKET_COMMAND_RPORTFWD_CLEAR  = 0x3
	SOCKET_COMMAND_RPORTFWD_REMOVE = 0x4

	SOCKET_COMMAND_SOCKSPROXY_ADD    = 0x5
	SOCKET_COMMAND_SOCKSPROXY_LIST   = 0x6
	SOCKET_COMMAND_SOCKSPROXY_REMOVE = 0x7
	SOCKET_COMMAND_SOCKSPROXY_CLEAR  = 0x8

	SOCKET_COMMAND_OPEN       = 0x10
	SOCKET_COMMAND_READ       = 0x11
	SOCKET_COMMAND_WRITE      = 0x12
	SOCKET_COMMAND_CLOSE      = 0x13
	SOCKET_COMMAND_CONNECT    = 0x14

	SOCKET_TYPE_REVERSE_PORTFWD = 0x1
	SOCKET_TYPE_REVERSE_PROXY   = 0x2
	SOCKET_TYPE_CLIENT          = 0x3

	SOCKET_ERROR_ALREADY_BOUND = 0x1
)

const (
	KERBEROS_COMMAND_LUID  = 0x0
	KERBEROS_COMMAND_KLIST = 0x1
	KERBEROS_COMMAND_PURGE = 0x2
	KERBEROS_COMMAND_PTT   = 0x3
)

const (
	COFFEELDR_FLAG_NON_THREADED = 0
	COFFEELDR_FLAG_THREADED     = 1
	COFFEELDR_FLAG_DEFAULT      = 2
)

const (
	INJECT_WAY_SPAWN   = 0
	INJECT_WAY_INJECT  = 1
	INJECT_WAY_EXECUTE = 2
)

const (
	THREAD_METHOD_DEFAULT            = 0
	THREAD_METHOD_CREATEREMOTETHREAD = 1
	THREAD_METHOD_NTCREATEHREADEX    = 2
	THREAD_METHOD_NTQUEUEAPCTHREAD   = 3
)

const (
	SecurityAnonymous      = 0x0
	SecurityIdentification = 0x1
	SecurityImpersonation  = 0x2
	SecurityDelegation     = 0x3
)

const (
	SECURITY_MANDATORY_UNTRUSTED_RID         = 0x00000000
	SECURITY_MANDATORY_LOW_RID               = 0x00001000
	SECURITY_MANDATORY_MEDIUM_RID            = 0x00002000
	SECURITY_MANDATORY_HIGH_RID              = 0x00003000
	SECURITY_MANDATORY_SYSTEM_RID            = 0x00004000
	SECURITY_MANDATORY_PROTECTED_PROCESS_RID = 0x00005000
)

const (
	TokenPrimary       = 1
	TokenImpersonation = 2
)

const (
	INJECT_ERROR_SUCCESS               = 0
	INJECT_ERROR_FAILED                = 1
	INJECT_ERROR_INVALID_PARAM         = 2
	INJECT_ERROR_PROCESS_ARCH_MISMATCH = 3
)
