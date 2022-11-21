// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: config.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileHashType int32

const (
	FileHashType_SHA256 FileHashType = 0
)

// Enum value maps for FileHashType.
var (
	FileHashType_name = map[int32]string{
		0: "SHA256",
	}
	FileHashType_value = map[string]int32{
		"SHA256": 0,
	}
)

func (x FileHashType) Enum() *FileHashType {
	p := new(FileHashType)
	*p = x
	return p
}

func (x FileHashType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileHashType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[0].Descriptor()
}

func (FileHashType) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[0]
}

func (x FileHashType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileHashType.Descriptor instead.
func (FileHashType) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

type Tool_Platform int32

const (
	Tool_PLATFORM_UNKNOWN Tool_Platform = 0
	Tool_PLATFORM_ANY     Tool_Platform = 1
	Tool_PLATFORM_LINUX   Tool_Platform = 2
	Tool_PLATFORM_WINDOWS Tool_Platform = 3
	Tool_PLATFORM_OSX     Tool_Platform = 4
)

// Enum value maps for Tool_Platform.
var (
	Tool_Platform_name = map[int32]string{
		0: "PLATFORM_UNKNOWN",
		1: "PLATFORM_ANY",
		2: "PLATFORM_LINUX",
		3: "PLATFORM_WINDOWS",
		4: "PLATFORM_OSX",
	}
	Tool_Platform_value = map[string]int32{
		"PLATFORM_UNKNOWN": 0,
		"PLATFORM_ANY":     1,
		"PLATFORM_LINUX":   2,
		"PLATFORM_WINDOWS": 3,
		"PLATFORM_OSX":     4,
	}
)

func (x Tool_Platform) Enum() *Tool_Platform {
	p := new(Tool_Platform)
	*p = x
	return p
}

func (x Tool_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tool_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[1].Descriptor()
}

func (Tool_Platform) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[1]
}

func (x Tool_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tool_Platform.Descriptor instead.
func (Tool_Platform) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1, 0}
}

type Tool_Architecture int32

const (
	Tool_ARCH_UNKNOWN Tool_Architecture = 0
	Tool_ARCH_ANY     Tool_Architecture = 1
	Tool_ARCH_X86     Tool_Architecture = 2
	Tool_ARCH_X64     Tool_Architecture = 3
	Tool_ARCH_ADM     Tool_Architecture = 4
)

// Enum value maps for Tool_Architecture.
var (
	Tool_Architecture_name = map[int32]string{
		0: "ARCH_UNKNOWN",
		1: "ARCH_ANY",
		2: "ARCH_X86",
		3: "ARCH_X64",
		4: "ARCH_ADM",
	}
	Tool_Architecture_value = map[string]int32{
		"ARCH_UNKNOWN": 0,
		"ARCH_ANY":     1,
		"ARCH_X86":     2,
		"ARCH_X64":     3,
		"ARCH_ADM":     4,
	}
)

func (x Tool_Architecture) Enum() *Tool_Architecture {
	p := new(Tool_Architecture)
	*p = x
	return p
}

func (x Tool_Architecture) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tool_Architecture) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[2].Descriptor()
}

func (Tool_Architecture) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[2]
}

func (x Tool_Architecture) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tool_Architecture.Descriptor instead.
func (Tool_Architecture) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1, 1}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of tools to serve
	Tool []*Tool `protobuf:"bytes,1,rep,name=tool,proto3" json:"tool,omitempty"`
	// Path to prepend to all other paths
	RootPath string `protobuf:"bytes,2,opt,name=root_path,json=rootPath,proto3" json:"root_path,omitempty"`
	// Hide the index page
	HideIndex bool `protobuf:"varint,3,opt,name=hide_index,json=hideIndex,proto3" json:"hide_index,omitempty"`
	// Listenaddr, as in 0.0.0.0:6666
	ListenAddr string `protobuf:"bytes,4,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	// Cache directory, may begin with ~/ for home directory.
	CacheDir string `protobuf:"bytes,5,opt,name=cache_dir,json=cacheDir,proto3" json:"cache_dir,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetTool() []*Tool {
	if x != nil {
		return x.Tool
	}
	return nil
}

func (x *Config) GetRootPath() string {
	if x != nil {
		return x.RootPath
	}
	return ""
}

func (x *Config) GetHideIndex() bool {
	if x != nil {
		return x.HideIndex
	}
	return false
}

func (x *Config) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *Config) GetCacheDir() string {
	if x != nil {
		return x.CacheDir
	}
	return ""
}

// Tool is a single servable tool entry
type Tool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Visible name for tool
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Path(s) under which to serve tool
	// Launching fails if this is not globally unique
	Path []string `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	// Sources to obtain the tool
	Source []*ToolSource `protobuf:"bytes,3,rep,name=source,proto3" json:"source,omitempty"`
	// Hashes to verify before serving
	Hash []*FileHash `protobuf:"bytes,4,rep,name=hash,proto3" json:"hash,omitempty"`
	// Should we attempt to embed this?
	Embed bool `protobuf:"varint,5,opt,name=embed,proto3" json:"embed,omitempty"`
	// Filename to provide when serving
	// If empty, last component of path is used
	Filename string `protobuf:"bytes,6,opt,name=filename,proto3" json:"filename,omitempty"`
	// Content-type, otherwise application/octet-stream
	ContentType string `protobuf:"bytes,7,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Platform and architecture
	Platform Tool_Platform     `protobuf:"varint,8,opt,name=platform,proto3,enum=config.Tool_Platform" json:"platform,omitempty"`
	Arch     Tool_Architecture `protobuf:"varint,9,opt,name=arch,proto3,enum=config.Tool_Architecture" json:"arch,omitempty"`
}

func (x *Tool) Reset() {
	*x = Tool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tool) ProtoMessage() {}

func (x *Tool) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tool.ProtoReflect.Descriptor instead.
func (*Tool) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Tool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tool) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Tool) GetSource() []*ToolSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Tool) GetHash() []*FileHash {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Tool) GetEmbed() bool {
	if x != nil {
		return x.Embed
	}
	return false
}

func (x *Tool) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Tool) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Tool) GetPlatform() Tool_Platform {
	if x != nil {
		return x.Platform
	}
	return Tool_PLATFORM_UNKNOWN
}

func (x *Tool) GetArch() Tool_Architecture {
	if x != nil {
		return x.Arch
	}
	return Tool_ARCH_UNKNOWN
}

// Toolsource describes a single source for getting the tool.
type ToolSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to obtain the tool.
	// May be a local path or a HTTP(S) path to be retrieved.
	SourcePath string `protobuf:"bytes,1,opt,name=source_path,json=sourcePath,proto3" json:"source_path,omitempty"`
	// If the source path is a supported archive, this path is extracted.
	ArchivePath string `protobuf:"bytes,2,opt,name=archive_path,json=archivePath,proto3" json:"archive_path,omitempty"`
	// If the file is compressed, but not an archive, decompress with this.
	DecompressWith string `protobuf:"bytes,3,opt,name=decompress_with,json=decompressWith,proto3" json:"decompress_with,omitempty"`
}

func (x *ToolSource) Reset() {
	*x = ToolSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolSource) ProtoMessage() {}

func (x *ToolSource) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolSource.ProtoReflect.Descriptor instead.
func (*ToolSource) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *ToolSource) GetSourcePath() string {
	if x != nil {
		return x.SourcePath
	}
	return ""
}

func (x *ToolSource) GetArchivePath() string {
	if x != nil {
		return x.ArchivePath
	}
	return ""
}

func (x *ToolSource) GetDecompressWith() string {
	if x != nil {
		return x.DecompressWith
	}
	return ""
}

type FileHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of hash to compute
	Type FileHashType `protobuf:"varint,1,opt,name=type,proto3,enum=config.FileHashType" json:"type,omitempty"`
	// Value of hash
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *FileHash) Reset() {
	*x = FileHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHash) ProtoMessage() {}

func (x *FileHash) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHash.ProtoReflect.Descriptor instead.
func (*FileHash) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *FileHash) GetType() FileHashType {
	if x != nil {
		return x.Type
	}
	return FileHashType_SHA256
}

func (x *FileHash) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa4, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x74,
	0x6f, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x44, 0x69, 0x72, 0x22, 0x81, 0x04,
	0x0a, 0x04, 0x54, 0x6f, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2d, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x54, 0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x22, 0x6e, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41,
	0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x02, 0x12,
	0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x53, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x4f, 0x53, 0x58, 0x10, 0x04, 0x22, 0x58, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x52, 0x43, 0x48, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43,
	0x48, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43, 0x48, 0x5f,
	0x58, 0x38, 0x36, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x58, 0x36,
	0x34, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x4d, 0x10,
	0x04, 0x22, 0x79, 0x0a, 0x0a, 0x54, 0x6f, 0x6f, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x57, 0x69, 0x74, 0x68, 0x22, 0x48, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x2a, 0x1a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36,
	0x10, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x61, 0x74, 0x69, 0x72, 0x2f, 0x72, 0x65, 0x64, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_proto_goTypes = []interface{}{
	(FileHashType)(0),      // 0: config.FileHashType
	(Tool_Platform)(0),     // 1: config.Tool.Platform
	(Tool_Architecture)(0), // 2: config.Tool.Architecture
	(*Config)(nil),         // 3: config.Config
	(*Tool)(nil),           // 4: config.Tool
	(*ToolSource)(nil),     // 5: config.ToolSource
	(*FileHash)(nil),       // 6: config.FileHash
}
var file_config_proto_depIdxs = []int32{
	4, // 0: config.Config.tool:type_name -> config.Tool
	5, // 1: config.Tool.source:type_name -> config.ToolSource
	6, // 2: config.Tool.hash:type_name -> config.FileHash
	1, // 3: config.Tool.platform:type_name -> config.Tool.Platform
	2, // 4: config.Tool.arch:type_name -> config.Tool.Architecture
	0, // 5: config.FileHash.type:type_name -> config.FileHashType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolSource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHash); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		EnumInfos:         file_config_proto_enumTypes,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
