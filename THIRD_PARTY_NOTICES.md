<!--
SPDX-License-Identifier: Apache-2.0
Copyright 2026 Hygon Information Technology Co., Ltd.
-->

# Third-Party Notices

本文件登记 `k8s-hcu-label-node` 通过 `go.mod` 引入的全部第三方 Go 模块依赖，包含直接依赖与间接依赖，共 42 项。

说明：

- 版本以 `go.mod` 中固定的版本号为准，校验和见 `go.sum`。
- 本仓库默认不提交 `vendor/` 目录（见 `.gitignore`）；执行 `go mod vendor` 后，各依赖即落盘于下文「本地路径」所示位置。未执行 vendor 时，依赖位于 `$GOMODCACHE/<module>@<version>`。
- 「版权声明」摘录自各依赖随包分发的 `LICENSE`/`NOTICE` 或源码文件头；各依赖原始许可证与版权文本以其包内文件为准，本文件仅作索引，不构成对原文的替代或改写。
- 所有依赖均按上游原样引入，HYGON 未做任何本地修改（no local modification / not forked / not patched）。
- 本文件不涵盖随仓库分发的 HCU 驱动运行库（`internal/pkg/lib/*.so`），该部分由 HCU Driver 提供，不属于 Go 模块依赖。

---

## 一、直接依赖（5 项）

### github.com/HYGON-AI/hcu-dcgm/v3
- 项目/仓库：https://github.com/HYGON-AI/hcu-dcgm
- 固定版本：v3.0.0
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/HYGON-AI/hcu-dcgm/v3
- 版权声明：Copyright (c) 2026 Hygon Information Technology Co., Ltd.; Copyright 2016-2019 by SW Group, Chengdu Haiguang IC Design Co., Ltd.; Copyright 2014 Advanced Micro Devices, Inc.
- HYGON 修改：HYGON原创
- 备注：该模块为 HYGON 自有开源组件；其源码树内另随包分发 HCU/AMD 驱动接口头文件（`pkg/dcgm/include/`，涉及 Apache-2.0、MIT、NCSA 三类许可证），详见该模块自带的 `THIRD_PARTY_NOTICES.md`。

### github.com/golang/glog
- 项目/仓库：https://github.com/golang/glog
- 固定版本：v1.2.2
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/golang/glog
- 版权声明：Copyright 2023 Google Inc. All Rights Reserved.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### k8s.io/api
- 项目/仓库：https://github.com/kubernetes/api
- 固定版本：v0.30.2
- 许可证：Apache-2.0
- 本地路径：vendor/k8s.io/api
- 版权声明：Copyright The Kubernetes Authors.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### k8s.io/apimachinery
- 项目/仓库：https://github.com/kubernetes/apimachinery
- 固定版本：v0.30.2
- 许可证：Apache-2.0
- 本地路径：vendor/k8s.io/apimachinery
- 版权声明：Copyright The Kubernetes Authors.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### k8s.io/client-go
- 项目/仓库：https://github.com/kubernetes/client-go
- 固定版本：v0.30.2
- 许可证：Apache-2.0
- 本地路径：vendor/k8s.io/client-go
- 版权声明：Copyright The Kubernetes Authors.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

---

## 二、间接依赖（37 项）

### github.com/davecgh/go-spew
- 项目/仓库：https://github.com/davecgh/go-spew
- 固定版本：v1.1.1
- 许可证：ISC
- 本地路径：vendor/github.com/davecgh/go-spew
- 版权声明：Copyright (c) 2012-2016 Dave Collins <dave@davec.name>; Permission to use, copy, modify, and/or distribute this software for any purpose with or without fee is hereby granted, provided that the above copyright notice and this permission notice appear in all copies.
- HYGON 修改：无（按上游原样引入）

### github.com/emicklei/go-restful/v3
- 项目/仓库：https://github.com/emicklei/go-restful
- 固定版本：v3.11.0
- 许可证：MIT
- 本地路径：vendor/github.com/emicklei/go-restful/v3
- 版权声明：Copyright (c) 2012,2013 Ernest Micklei; The above copyright notice and this permission notice shall be included in all copies; NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
- HYGON 修改：无（按上游原样引入）

### github.com/go-logr/logr
- 项目/仓库：https://github.com/go-logr/logr
- 固定版本：v1.4.1
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/go-logr/logr
- 版权声明：Copyright 2021 The logr Authors.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### github.com/go-openapi/jsonpointer
- 项目/仓库：https://github.com/go-openapi/jsonpointer
- 固定版本：v0.21.0
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/go-openapi/jsonpointer
- 版权声明：Copyright 2013 sigu-399 ( https://github.com/sigu-399 ); Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### github.com/go-openapi/jsonreference
- 项目/仓库：https://github.com/go-openapi/jsonreference
- 固定版本：v0.21.0
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/go-openapi/jsonreference
- 版权声明：Copyright 2013 sigu-399 ( https://github.com/sigu-399 ); Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### github.com/go-openapi/swag
- 项目/仓库：https://github.com/go-openapi/swag
- 固定版本：v0.23.0
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/go-openapi/swag
- 版权声明：Copyright 2015 go-swagger maintainers; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### github.com/gogo/protobuf
- 项目/仓库：https://github.com/gogo/protobuf
- 固定版本：v1.3.2
- 许可证：BSD-3-Clause
- 本地路径：vendor/github.com/gogo/protobuf
- 版权声明：Copyright (c) 2013, The GoGo Authors. All rights reserved.; Copyright 2010 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `AUTHORS`、`CONTRIBUTORS` 文件，列明贡献者归属。

### github.com/golang/protobuf
- 项目/仓库：https://github.com/golang/protobuf
- 固定版本：v1.5.4
- 许可证：BSD-3-Clause
- 本地路径：vendor/github.com/golang/protobuf
- 版权声明：Copyright 2010 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `AUTHORS`、`CONTRIBUTORS` 文件，列明贡献者归属。

### github.com/google/gnostic-models
- 项目/仓库：https://github.com/google/gnostic-models
- 固定版本：v0.6.8
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/google/gnostic-models
- 版权声明：Copyright 2017 Google LLC. All Rights Reserved.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### github.com/google/gofuzz
- 项目/仓库：https://github.com/google/gofuzz
- 固定版本：v1.2.0
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/google/gofuzz
- 版权声明：Copyright 2014 Google Inc. All rights reserved.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### github.com/google/uuid
- 项目/仓库：https://github.com/google/uuid
- 固定版本：v1.6.0
- 许可证：BSD-3-Clause
- 本地路径：vendor/github.com/google/uuid
- 版权声明：Copyright (c) 2009,2014 Google Inc. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `CONTRIBUTORS` 文件，列明贡献者归属。

### github.com/imdario/mergo
- 项目/仓库：https://github.com/imdario/mergo
- 固定版本：v0.3.6
- 许可证：BSD-3-Clause
- 本地路径：vendor/github.com/imdario/mergo
- 版权声明：Copyright (c) 2013 Dario Castañé. All rights reserved.; Copyright (c) 2012 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）

### github.com/josharian/intern
- 项目/仓库：https://github.com/josharian/intern
- 固定版本：v1.0.0
- 许可证：MIT
- 本地路径：vendor/github.com/josharian/intern
- 版权声明：Copyright (c) 2019 Josh Bleecher Snyder; The above copyright notice and this permission notice shall be included in all copies; AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
- HYGON 修改：无（按上游原样引入）
- 备注：许可证文件名为 `license.md`。

### github.com/json-iterator/go
- 项目/仓库：https://github.com/json-iterator/go
- 固定版本：v1.1.12
- 许可证：MIT
- 本地路径：vendor/github.com/json-iterator/go
- 版权声明：Copyright (c) 2016 json-iterator; The above copyright notice and this permission notice shall be included in all copies; AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
- HYGON 修改：无（按上游原样引入）

### github.com/mailru/easyjson
- 项目/仓库：https://github.com/mailru/easyjson
- 固定版本：v0.7.7
- 许可证：MIT
- 本地路径：vendor/github.com/mailru/easyjson
- 版权声明：Copyright (c) 2016 Mail.Ru Group; The above copyright notice and this permission notice shall be included in all copies; 部分文件含 Copyright (c) 2009 The Go Authors. All rights reserved.
- HYGON 修改：无（按上游原样引入）

### github.com/modern-go/concurrent
- 项目/仓库：https://github.com/modern-go/concurrent
- 固定版本：v0.0.0-20180306012644-bacd9c7ef1dd
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/modern-go/concurrent
- 版权声明：许可证文件为 Apache License 2.0 标准文本，未附加独立版权行；归属见上游仓库 modern-go 组织。
- HYGON 修改：无（按上游原样引入）

### github.com/modern-go/reflect2
- 项目/仓库：https://github.com/modern-go/reflect2
- 固定版本：v1.0.2
- 许可证：Apache-2.0
- 本地路径：vendor/github.com/modern-go/reflect2
- 版权声明：许可证文件为 Apache License 2.0 标准文本，未附加独立版权行；归属见上游仓库 modern-go 组织。
- HYGON 修改：无（按上游原样引入）

### github.com/munnerz/goautoneg
- 项目/仓库：https://github.com/munnerz/goautoneg
- 固定版本：v0.0.0-20191010083416-a7dc8b61c822
- 许可证：BSD-3-Clause
- 本地路径：vendor/github.com/munnerz/goautoneg
- 版权声明：Copyright (c) 2011, Open Knowledge Foundation Ltd. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）

### github.com/spf13/pflag
- 项目/仓库：https://github.com/spf13/pflag
- 固定版本：v1.0.5
- 许可证：BSD-3-Clause
- 本地路径：vendor/github.com/spf13/pflag
- 版权声明：Copyright (c) 2012 Alex Ogier. All rights reserved.; Copyright (c) 2012 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）

### go.etcd.io/bbolt
- 项目/仓库：https://github.com/etcd-io/bbolt
- 固定版本：v1.3.11
- 许可证：MIT
- 本地路径：vendor/go.etcd.io/bbolt
- 版权声明：Copyright (c) 2013 Ben Johnson; The above copyright notice and this permission notice shall be included in all copies; AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
- HYGON 修改：无（按上游原样引入）

### golang.org/x/net
- 项目/仓库：https://cs.opensource.google/go/x/net （镜像：https://github.com/golang/net）
- 固定版本：v0.28.0
- 许可证：BSD-3-Clause
- 本地路径：vendor/golang.org/x/net
- 版权声明：Copyright 2009 The Go Authors.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `PATENTS` 文件（Go 专利授权声明），须随源码一并保留。

### golang.org/x/oauth2
- 项目/仓库：https://cs.opensource.google/go/x/oauth2 （镜像：https://github.com/golang/oauth2）
- 固定版本：v0.18.0
- 许可证：BSD-3-Clause
- 本地路径：vendor/golang.org/x/oauth2
- 版权声明：Copyright (c) 2009 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）

### golang.org/x/sys
- 项目/仓库：https://cs.opensource.google/go/x/sys （镜像：https://github.com/golang/sys）
- 固定版本：v0.24.0
- 许可证：BSD-3-Clause
- 本地路径：vendor/golang.org/x/sys
- 版权声明：Copyright 2009 The Go Authors.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `PATENTS` 文件（Go 专利授权声明），须随源码一并保留。

### golang.org/x/term
- 项目/仓库：https://cs.opensource.google/go/x/term （镜像：https://github.com/golang/term）
- 固定版本：v0.23.0
- 许可证：BSD-3-Clause
- 本地路径：vendor/golang.org/x/term
- 版权声明：Copyright 2009 The Go Authors.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `PATENTS` 文件（Go 专利授权声明），须随源码一并保留。

### golang.org/x/text
- 项目/仓库：https://cs.opensource.google/go/x/text （镜像：https://github.com/golang/text）
- 固定版本：v0.17.0
- 许可证：BSD-3-Clause
- 本地路径：vendor/golang.org/x/text
- 版权声明：Copyright 2009 The Go Authors.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `PATENTS` 文件（Go 专利授权声明），须随源码一并保留。

### golang.org/x/time
- 项目/仓库：https://cs.opensource.google/go/x/time （镜像：https://github.com/golang/time）
- 固定版本：v0.3.0
- 许可证：BSD-3-Clause
- 本地路径：vendor/golang.org/x/time
- 版权声明：Copyright (c) 2009 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `PATENTS` 文件（Go 专利授权声明），须随源码一并保留。

### google.golang.org/appengine
- 项目/仓库：https://github.com/golang/appengine
- 固定版本：v1.6.8
- 许可证：Apache-2.0
- 本地路径：vendor/google.golang.org/appengine
- 版权声明：Copyright 2011 Google Inc. All rights reserved.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### google.golang.org/protobuf
- 项目/仓库：https://github.com/protocolbuffers/protobuf-go
- 固定版本：v1.34.2
- 许可证：BSD-3-Clause
- 本地路径：vendor/google.golang.org/protobuf
- 版权声明：Copyright (c) 2018 The Go Authors. All rights reserved.; 部分文件含 Copyright 2008 Google Inc. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：包内另附 `PATENTS` 文件（Go 专利授权声明），须随源码一并保留。

### gopkg.in/inf.v0
- 项目/仓库：https://github.com/go-inf/inf
- 固定版本：v0.9.1
- 许可证：BSD-3-Clause
- 本地路径：vendor/gopkg.in/inf.v0
- 版权声明：Copyright (c) 2012 Péter Surányi. Portions Copyright (c) 2009 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）

### gopkg.in/yaml.v2
- 项目/仓库：https://github.com/go-yaml/yaml （分支 v2）
- 固定版本：v2.4.0
- 许可证：Apache-2.0 与 MIT 双重（Go 代码为 Apache-2.0；源自 libyaml 的移植部分为 MIT）
- 本地路径：vendor/gopkg.in/yaml.v2
- 版权声明：Copyright 2011-2016 Canonical Ltd. (NOTICE, Apache-2.0); Copyright (c) 2006 Kirill Simonov (LICENSE.libyaml, MIT); The above copyright notice and this permission notice shall be included in all copies
- HYGON 修改：无（按上游原样引入）
- 备注：包内同时提供 `LICENSE`、`LICENSE.libyaml`、`NOTICE` 三个文件，分发时须一并保留。

### gopkg.in/yaml.v3
- 项目/仓库：https://github.com/go-yaml/yaml （分支 v3）
- 固定版本：v3.0.1
- 许可证：Apache-2.0 与 MIT 双重（Go 代码为 Apache-2.0；源自 libyaml 的移植部分为 MIT）
- 本地路径：vendor/gopkg.in/yaml.v3
- 版权声明：Copyright 2011-2016 Canonical Ltd. (NOTICE, Apache-2.0); Copyright (c) 2011-2019 Canonical Ltd 与 Copyright (c) 2006-2010 Kirill Simonov (MIT 部分); The above copyright notice and this permission notice shall be included in all copies
- HYGON 修改：无（按上游原样引入）
- 备注：包内同时提供 `LICENSE`、`NOTICE` 两个文件，分发时须一并保留。

### k8s.io/klog/v2
- 项目/仓库：https://github.com/kubernetes/klog
- 固定版本：v2.120.1
- 许可证：Apache-2.0
- 本地路径：vendor/k8s.io/klog/v2
- 版权声明：Copyright 2021 The Kubernetes Authors.; 部分文件含 Copyright 2013 Google Inc. All Rights Reserved.（源自 glog 的 fork）; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### k8s.io/kube-openapi
- 项目/仓库：https://github.com/kubernetes/kube-openapi
- 固定版本：v0.0.0-20240228011516-70dd3763d340
- 许可证：Apache-2.0
- 本地路径：vendor/k8s.io/kube-openapi
- 版权声明：Copyright The Kubernetes Authors.; 部分文件含 Copyright 2015 go-swagger maintainers 与 Copyright 2020 The Go Authors. All rights reserved.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### k8s.io/utils
- 项目/仓库：https://github.com/kubernetes/utils
- 固定版本：v0.0.0-20230726121419-3b25d923346b
- 许可证：Apache-2.0
- 本地路径：vendor/k8s.io/utils
- 版权声明：Copyright The Kubernetes Authors.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### sigs.k8s.io/json
- 项目/仓库：https://github.com/kubernetes-sigs/json
- 固定版本：v0.0.0-20221116044647-bc3834ca7abd
- 许可证：Apache-2.0 与 BSD-3-Clause 双重（自有代码为 Apache-2.0；源自 Go 标准库 encoding/json 的部分为 BSD-3-Clause）
- 本地路径：vendor/sigs.k8s.io/json
- 版权声明：Copyright 2021 The Kubernetes Authors.; Copyright (c) 2009 The Go Authors. All rights reserved.; Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer
- HYGON 修改：无（按上游原样引入）
- 备注：`LICENSE` 文件在 Apache-2.0 正文之后附加了 Go 标准库的 BSD-3-Clause 全文，两段均须保留。

### sigs.k8s.io/structured-merge-diff/v4
- 项目/仓库：https://github.com/kubernetes-sigs/structured-merge-diff
- 固定版本：v4.4.1
- 许可证：Apache-2.0
- 本地路径：vendor/sigs.k8s.io/structured-merge-diff/v4
- 版权声明：Copyright 2018 The Kubernetes Authors.; Licensed under the Apache License, Version 2.0
- HYGON 修改：无（按上游原样引入）

### sigs.k8s.io/yaml
- 项目/仓库：https://github.com/kubernetes-sigs/yaml
- 固定版本：v1.3.0
- 许可证：MIT 与 BSD-3-Clause 双重（自有代码为 MIT；源自 Go 标准库的部分为 BSD-3-Clause）
- 本地路径：vendor/sigs.k8s.io/yaml
- 版权声明：Copyright (c) 2014 Sam Ghods (MIT); Copyright (c) 2012 The Go Authors. All rights reserved. (BSD-3-Clause); The above copyright notice and this permission notice shall be included in all copies
- HYGON 修改：无（按上游原样引入）
- 备注：`LICENSE` 文件同时包含 MIT 与 BSD-3-Clause 两段全文，分发时须一并保留。

---

## 三、许可证汇总

| 许可证 | 依赖数 | 依赖 |
| --- | --- | --- |
| Apache-2.0 | 21（其中 3 项为双重许可） | hcu-dcgm/v3、glog、k8s.io/api、k8s.io/apimachinery、k8s.io/client-go、go-logr/logr、go-openapi/jsonpointer、go-openapi/jsonreference、go-openapi/swag、gnostic-models、gofuzz、modern-go/concurrent、modern-go/reflect2、appengine、klog/v2、kube-openapi、k8s.io/utils、structured-merge-diff/v4、yaml.v2\*、yaml.v3\*、sigs.k8s.io/json\* |
| BSD-3-Clause | 16（其中 2 项为双重许可） | gogo/protobuf、golang/protobuf、google/uuid、imdario/mergo、munnerz/goautoneg、spf13/pflag、x/net、x/oauth2、x/sys、x/term、x/text、x/time、google.golang.org/protobuf、gopkg.in/inf.v0、sigs.k8s.io/json\*、sigs.k8s.io/yaml\* |
| MIT | 8（其中 3 项为双重许可） | go-restful/v3、josharian/intern、json-iterator/go、mailru/easyjson、go.etcd.io/bbolt、yaml.v2\*、yaml.v3\*、sigs.k8s.io/yaml\* |
| ISC | 1 | davecgh/go-spew |

标注 `*` 者为双重许可，在对应的多行中重复计入。按模块去重后共 42 项：单一 Apache-2.0 18 项、单一 BSD-3-Clause 14 项、单一 MIT 5 项、ISC 1 项、双重许可 4 项（yaml.v2、yaml.v3、sigs.k8s.io/json、sigs.k8s.io/yaml）。

合规结论：上述许可证均为宽松型（permissive），与本项目 Apache-2.0 许可证兼容，不含 GPL/LGPL/AGPL 等 copyleft 传染性许可证。Apache-2.0、BSD-3-Clause、MIT、ISC 均要求在二进制或源码分发时保留原始版权与许可证声明，本文件及各依赖包内的 `LICENSE`/`NOTICE`/`PATENTS` 文件共同满足该义务。

## 四、维护说明

依赖变更（新增、删除、升版）后须同步更新本文件。可用如下命令核对当前依赖清单与版本：

```bash
go list -m all
go mod vendor && ls vendor/
```
