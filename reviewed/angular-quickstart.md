
<!-- TOC -->

- [Angular Quickstart](#angular-quickstart)
  - [Environment](#environment)
    - [Visual Studio Code](#visual-studio-code)
    - [Angular CLI](#angular-cli)
  - [Introduction Visual Studio Code](#introduction-visual-studio-code)
    - [Extensions](#extensions)
  - [Introduction Angular CLI](#introduction-angular-cli)
    - [Command](#command)
      - [Usage](#usage)
      - [Generating](#generating)
    - [Flow](#flow)
- [Reference](#reference)

<!-- /TOC -->

# Angular Quickstart


## Environment


### Visual Studio Code

*  [Download from Visual Studio Code Website](https://code.visualstudio.com/download)

### Angular CLI

1. Install Node.js at least node `6.9.x` and npm `3.x.x`
    * [Downlaod from Node.js Website](https://nodejs.org/en/download/)
    * [nvm](https://github.com/creationix/nvm)
    * [n](https://github.com/tj/n)
2. Install Angular CLI
    * `npm install -g @angular/cli`


## Introduction Visual Studio Code

* [Visual Studio Code Basic Editing](https://code.visualstudio.com/docs/editor/codebasics)
* [Visual Studio Code 快速上手指南 (20160530)](https://www.slideshare.net/shengyou/visual-studio-code-62532711)
* [活用 Visual Studio Code (20150919)](https://channel9.msdn.com/Series/Mastering-Visual-Studio-Code)

Capture from [Visual Studio Code Tips and Tricks](https://github.com/Microsoft/vscode-tips-and-tricks):

* Basics
  * [Command Palette](https://github.com/Microsoft/vscode-tips-and-tricks#command-palette)
  * [Quick open](https://github.com/Microsoft/vscode-tips-and-tricks#quick-open)
  * [CLI tool](https://github.com/Microsoft/vscode-tips-and-tricks#cli-tool)
  * [.vscode folder](https://github.com/Microsoft/vscode-tips-and-tricks#vscode-folder)
  * [Status bar decorations](https://github.com/Microsoft/vscode-tips-and-tricks#status-bar-decorations)
* Customization
  * [Theme](https://github.com/Microsoft/vscode-tips-and-tricks#change-your-theme)
  * [Keyboard Shortcuts](https://github.com/Microsoft/vscode-tips-and-tricks#change-your-keyboard-shortcuts)
  * [Settings](https://github.com/Microsoft/vscode-tips-and-tricks#tune-your-settings)
* [Extensions](https://github.com/Microsoft/vscode-tips-and-tricks#extensions)
* [File and folder management](https://github.com/Microsoft/vscode-tips-and-tricks#file-and-folder-management)
* [Editing hacks](https://github.com/Microsoft/vscode-tips-and-tricks#editing-hacks)
* [Intellisense](https://github.com/Microsoft/vscode-tips-and-tricks#intellisense)
* [Snippets](https://github.com/Microsoft/vscode-tips-and-tricks#snippets)
  * `File -> Preferences -> User Snippets`
* [Git Integration](https://github.com/Microsoft/vscode-tips-and-tricks#git-integration)
* [Debugging](https://github.com/Microsoft/vscode-tips-and-tricks#debugging)
* [Task Runner](https://github.com/Microsoft/vscode-tips-and-tricks#task-runner)


### Extensions

* Recommended Extensions
  * [開發 Angular 必備 Visual Studio Code 擴充套件](https://paper.dropbox.com/doc/Angular-VSCode--Kh2w3saOyZtJSHawFoBem)
  * [Angular Extension Pack for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=doggy8088.angular-extension-pack)
* Install
  * [Extension Marketplace](https://code.visualstudio.com/docs/editor/extension-gallery#_browse-and-install-extensions)
  * [Extension Packs](https://code.visualstudio.com/docs/extensionAPI/extension-manifest#_extension-packs)
  * [Workspace Recommended Extensions](https://code.visualstudio.com/docs/editor/extension-gallery#_workspace-recommended-extensions)
  * [Command line extension management](https://code.visualstudio.com/docs/editor/extension-gallery#_command-line-extension-management)


## Introduction Angular CLI

* `The Angular CLI is a command line interface tool that can create a project, add files, and perform a variety of ongoing development tasks such as testing, bundling, and deployment.`
* `The goal is to build and run a simple Angular application in TypeScript, using the Angular CLI while adhering to the Style Guide recommendations that benefit every Angular project.`


### Command

* [Angular CLI Wiki](https://github.com/angular/angular-cli/wiki)

#### Usage

```bash
ng help
ng set -g packageManager=yarn
ng new PROJECT-NAME
cd PROJECT-NAME
ng serve
```

#### Generating

Scaffold  | Usage
---       | ---
[Component](https://github.com/angular/angular-cli/wiki/generate-component) | `ng g component my-new-component`
[Directive](https://github.com/angular/angular-cli/wiki/generate-directive) | `ng g directive my-new-directive`
[Pipe](https://github.com/angular/angular-cli/wiki/generate-pipe)           | `ng g pipe my-new-pipe`
[Service](https://github.com/angular/angular-cli/wiki/generate-service)     | `ng g service my-new-service`
[Class](https://github.com/angular/angular-cli/wiki/generate-class)         | `ng g class my-new-class`
[Guard](https://github.com/angular/angular-cli/wiki/generate-guard)         | `ng g guard my-new-guard`
[Interface](https://github.com/angular/angular-cli/wiki/generate-interface) | `ng g interface my-new-interface`
[Enum](https://github.com/angular/angular-cli/wiki/generate-enum)           | `ng g enum my-new-enum`
[Module](https://github.com/angular/angular-cli/wiki/generate-module)       | `ng g module my-module`


### Flow

[Todo Example Project](https://github.com/7Hd/angular-bootstrap-todo):

* `ng new todo --skip-tests --prefix td`
* `ng generate component todo-list --inline-style`
* `ng generate component todo --inline-style --inline-template --flat`
* `ng generate module shared --module app`
* `ng generate service shared\todo-data --module shared`
* `ng generate class models\todo`
* `ng generate enum shared\todo-priority`
* `ng generate pipe shared\enum-array --export shared`


# Reference

* [VS Code Official Document](https://code.visualstudio.com/docs)
* [Angular Official CLI Quickstart](https://angular.io/docs/ts/latest/cli-quickstart.html)
* [taobao fed: nvm or n (2015-11-17)](http://taobaofed.org/blog/2015/11/17/nvm-or-n/)
* [doggy8088/Angular4開發環境說明.md](https://gist.github.com/doggy8088/15e434b43992cf25a78700438743774a)
* [Youtube Editor 編輯者](https://www.youtube.com/channel/UC8-c0VKKqkG_aPe0RG3SF0A)
* [Angular Taiwan](https://forum.angular.tw)
* [Awesome VSCode](https://github.com/viatsko/awesome-vscode)
