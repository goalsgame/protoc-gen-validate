package ccnop

const headerFileTpl = `// Code generated by protoc-gen-validate
// source: {{ .InputPath }}
// DO NOT EDIT!!!

#pragma once

#include <string>

#include "validate/pgvvalidate.h"
#include "{{ output .File ".h" }}"

{{ range .Package.ProtoName.SplitOnDot }}
namespace {{ . }} {
{{- end }}

using std::string;

{{ range .AllMessages }}
	extern inline bool Validate(__attribute__((unused)) const {{ class . }}& m, __attribute__((unused)) pgv::ValidationMsg* err) { return true; }
{{ end }}

{{ range .Package.ProtoName.SplitOnDot -}}
} // namespace
{{ end }}

{{ range .AllMessages -}}
{{- if not (ignored .) -}}
{{ end -}}
{{ end }}
`

const moduleFileTpl = `// Code generated by protoc-gen-validate
// source: {{ .InputPath }}
// DO NOT EDIT!!!

namespace pgv {
namespace validate {

} // namespace validate
} // namespace pgv
`
