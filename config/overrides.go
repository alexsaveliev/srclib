// +build off
// TODO(sqs): reenable

package config

import (
	"sourcegraph.com/sourcegraph/srclib/repo"
	"sourcegraph.com/sourcegraph/srclib/toolchain/javascript"
	"sourcegraph.com/sourcegraph/srclib/toolchain/python"
	"sourcegraph.com/sourcegraph/srclib/toolchain/ruby"
	"sourcegraph.com/sourcegraph/srclib/unit"
)

var overrides = map[repo.URI]*Repository{
	"github.com/emicklei/go-restful": &Repository{
		ScanIgnore: []string{"./examples"},
	},
	"github.com/ruby/ruby": &Repository{
		SourceUnits: unit.SourceUnits{
			&ruby.RubyLib{
				LibName: ".",
				Dir:     ".",
				Files: []string{
					"addr2line.c",
					"array.c",
					"bignum.c",
					"class.c",
					"compar.c",
					"compile.c",
					"complex.c",
					"cont.c",
					"debug.c",
					"dir.c",
					"dln.c",
					"dln_find.c",
					"dmydln.c",
					"dmyext.c",
					"encoding.c",
					"enum.c",
					"enumerator.c",
					"error.c",
					"eval.c",
					"eval_error.c",
					"eval_jump.c",
					"file.c",
					"gc.c",
					"goruby.c",
					"hash.c",
					"inits.c",
					"io.c",
					"iseq.c",
					"load.c",
					"loadpath.c",
					"localeinit.c",
					"main.c",
					"marshal.c",
					"math.c",
					"miniinit.c",
					"node.c",
					"numeric.c",
					"object.c",
					"pack.c",
					"proc.c",
					"process.c",
					"random.c",
					"range.c",
					"rational.c",
					"re.c",
					"regcomp.c",
					"regenc.c",
					"regerror.c",
					"regexec.c",
					"regparse.c",
					"regsyntax.c",
					"ruby.c",
					"safe.c",
					"signal.c",
					"siphash.c",
					"sparc.c",
					"sprintf.c",
					"st.c",
					"strftime.c",
					"string.c",
					"struct.c",
					"thread.c",
					"thread_pthread.c",
					"thread_win32.c",
					"time.c",
					"transcode.c",
					"util.c",
					"variable.c",
					"version.c",
					"vm_backtrace.c",
					"vm.c",
					"vm_dump.c",
					"vm_eval.c",
					"vm_exec.c",
					"vm_insnhelper.c",
					"vm_method.c",
					"vm_trace.c",
					"vsnprintf.c",
					"lib/base64.rb",
					"lib/benchmark.rb",
					"lib/complex.rb",
					"lib/csv.rb",
					"lib/English.rb",
					"lib/erb.rb",
					"lib/fileutils.rb",
					"lib/find.rb",
					"lib/ipaddr.rb",
					"lib/logger.rb",
					"lib/mathn.rb",
					"lib/net/ftp.rb",
					"lib/net/http/backward.rb",
					"lib/net/http/exceptions.rb",
					"lib/net/http/generic_request.rb",
					"lib/net/http/header.rb",
					"lib/net/http/proxy_delta.rb",
					"lib/net/http.rb",
					"lib/net/http/request.rb",
					"lib/net/http/requests.rb",
					"lib/net/http/response.rb",
					"lib/net/http/responses.rb",
					"lib/net/https.rb",
					"lib/net/imap.rb",
					"lib/net/pop.rb",
					"lib/net/protocol.rb",
					"lib/net/smtp.rb",
					"lib/net/telnet.rb",
					"lib/ostruct.rb",
					"lib/pp.rb",
					"lib/prettyprint.rb",
					"lib/prime.rb",
					"lib/profile.rb",
					"lib/profiler.rb",
					"lib/rational.rb",
					"lib/scanf.rb",
					"lib/securerandom.rb",
					"lib/set.rb",
					"lib/shellwords.rb",
					"lib/singleton.rb",
					"lib/sync.rb",
					"lib/tempfile.rb",
					"lib/timeout.rb",
					"lib/time.rb",
					"lib/tmpdir.rb",
					"lib/uri/common.rb",
					"lib/uri/ftp.rb",
					"lib/uri/generic.rb",
					"lib/uri/http.rb",
					"lib/uri/https.rb",
					"lib/uri/ldap.rb",
					"lib/uri/ldaps.rb",
					"lib/uri/mailto.rb",
					"lib/uri.rb",
					"ext/json",
					"ext/date",
					"ext/io",
					"ext/readline",
					"ext/pathname",
					"ext/digest",
					"ext/bigdecimal",
				},
			},
		},
		ScanIgnore: []string{"./ext", "./test"},
		Global: Global{
			"ruby": &ruby.Config{OmitStdlib: true},
		},
	},
	"code.google.com/p/go": &Repository{
		ScanIgnore: []string{"./misc", "./test", "./doc", "./cmd", "./src/cmd"},
		// TODO(sqs): set base import path == "" for stdlib?
	},
	"github.com/joyent/node": &Repository{
		SourceUnits: unit.SourceUnits{
			&javascript.CommonJSPackage{
				Package:            []byte(`{}`),
				PackageName:        javascript.NodeJSStdlibUnit,
				PackageDescription: "The Node.js core API.",
				Dir:                ".",
				LibFiles: []string{
					"lib/assert.js",
					"lib/buffer.js",
					"lib/child_process.js",
					"lib/cluster.js",
					"lib/console.js",
					"lib/constants.js",
					"lib/crypto.js",
					"lib/dgram.js",
					"lib/dns.js",
					"lib/domain.js",
					"lib/events.js",
					"lib/freelist.js",
					"lib/fs.js",
					"lib/http.js",
					"lib/https.js",
					"lib/module.js",
					"lib/net.js",
					"lib/os.js",
					"lib/path.js",
					"lib/punycode.js",
					"lib/querystring.js",
					"lib/readline.js",
					"lib/repl.js",
					"lib/smalloc.js",
					"lib/stream.js",
					"lib/string_decoder.js",
					"lib/sys.js",
					"lib/timers.js",
					"lib/tls.js",
					"lib/tty.js",
					"lib/url.js",
					"lib/util.js",
					"lib/vm.js",
					"lib/zlib.js",
				},
			},
		},

		// Suppress the Python source unit that exists because the node
		// repo has *.py files.
		ScanIgnoreUnitTypes: []string{unit.Type(&python.DistPackage{})},
		ScanIgnore:          []string{"./tools", "./deps", "./test", "./src"},

		Global: Global{
			"jsg": &javascript.JSGConfig{
				Plugins: map[string]interface{}{
					// In this repository, the node core modules are in the
					// lib/ dir.
					"node": map[string]string{"coreModulesDir": "lib/"},

					"$(JSG_DIR)/node_modules/tern-node-api-doc/node-api-doc": map[string]string{
						"apiDocDir":      "doc/api/",
						"apiSrcDir":      "lib/",
						"generateJSPath": "tools/doc/generate.js",
					},
				},
			},
		},
	},
}
