# win32_sfx
make self extracting file run on Windows


## 资料 

- [How to create Windows compatible rar sfx archive on Linux] https://stackoverflow.com/questions/30479901/how-to-create-windows-compatible-rar-sfx-archive-on-linux
- 描述SFX自解压运行后执行的指令 https://github.com/osmc/osmc/blob/master/installer/host/qt_host_installer/winrar.sfx
- SFX Script Command 含义 https://msfn.org/board/topic/34343-winrar-sfx-commands/
- [与想象的步骤不一样，未验证生成产物] Golang 版本的SFX 自解压 https://github.com/touchifyapp/sfx
- 7-ZIP SFX 格式 https://blog.csdn.net/hxbb00/article/details/79055964


## 借助WinRAR等价实现
```
=> pack to sfx
<winrar.exe> -y a -afzip -ep1 -sfxzip.sfx <dst file path>
  <files to pack> -z"<comment file path>" -iicon"<icon file path>"
  
-y -> 所有回答都是 yes
a 向压缩包内添加文件
-afzip zip格式。 rar.exe 只支持 RAR 格式，winrar.exe 支持 RAR和ZIP 格式
-ep1 是压缩包内路径的处理方式，响应的有 ep2 ep 等
-z 是压缩包的注释

=> change name in zip
<winrar.exe> rn <dst file path> <src name in zip> <dst name in zip>
```

## 有了 SFX 文件，你可能需要修改 PE 文件的能力

- [C++] https://github.com/TACIXAT/portable-executable-library
- [C++] https://github.com/trailofbits/pe-parse
- [python][需要跟 MSVC 协作使用，不算跨平台] https://github.com/avast/pe_tools
- [python][未验证] https://github.com/kd8bny/pedit
- [python][只能打印一堆信息]https://github.com/erocarrera/pefile

- [golang][需要适配golang 1.12][没找到感兴趣的]https://github.com/soluwalana/pefile-go
