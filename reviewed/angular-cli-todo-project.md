
<!-- TOC -->

- [Angular CLI Todo Project](#angular-cli-todo-project)
  - [Create Project](#create-project)
  - [Basic CRUD](#basic-crud)
    - [Todo Model](#todo-model)
    - [Todo List](#todo-list)
    - [Create](#create)
    - [Read](#read)
    - [Update](#update)
    - [Delete](#delete)
  - [Other](#other)

<!-- /TOC -->

# Angular CLI Todo Project

## Create Project

1. 利用 Angular CLI 建立專案: `ng new todomvc-angular-cli`
2. 複製 `index.html`, `assets/css/index.css` 到 CLI 建立的專案內
    * `index.html`: `src/app/app.component.html`
      * 只需要複製 `<body>` 內的部分即可
    * `assets/css/index.css`: `src/styles.css`
3. 測試修改結果
    * `npm start`: 與 `ng serve` 相同
    * `npm run build`: 與 `ng build` 相同
      * 直接開啟編譯結果須增加 `--base-href .`


## Basic CRUD

### Todo Model

1. 建立 Todo Model 的 `Class` 或 `interface`
    * 手動新增檔案
    * 透過Cli: `ng generate class [PATH/]NAME`
      * 等同於 `ng g cl [PATH/]NAME`
2. 修改 Model 內容
    * 包含 `id`, `status`, `name` 3 項屬性
    ```ts
    export class Todo {
      id?: number;
      status: TodoStatus;
      name: string;

      constructor(obj: any = {}) {
        ...
      }
    }
    ```
3. `status` 使用 `Enum`
    * 手動建立
    * CLI: `ng generate enum [PATH/]NAME`


### Todo List

1. 於 `app.component.ts` 內增加
    ```ts
    export class AppComponent {
      todoList: Todo[] = [];
    }
    ```
2. 建立初始值
    ```ts
    todoList: Todo[] = [
      new Todo({...}),
    ];
    ```


### Create

1. 修改 `app.component.html`
    * 原始
      ```html
      <input class="new-todo" placeholder="What needs to be done?" autofocus>
      ```
    * 修改後
      ```html
      <input #newTodo (keyup.enter)="createNewTodo(newTodo)" class="new-todo" placeholder="What needs to be done?" autofocus>
      ```
      * [Template reference variables](https://angular.io/guide/template-syntax#template-reference-variables--var-)
      * [Event binding](https://angular.io/guide/template-syntax#event-binding)
2. 增加對應 `Function`
    ```ts
    createNewTodo(input: HTMLInputElement) {
      this.todoList.push(new Todo({...}));
      input.value = '';
    }
    ```
      * 新增 todo 並清除畫面的輸入欄位


### Read

1. 利用 [`ngFor`](https://angular.io/guide/template-syntax#ngfor) 來顯示
    ```html
    <ul class="todo-list" *ngFor="let todo of todoList">
      <li>
        ...
        <label>{{todo.name}}</label>
      </li>
    </ul>
    ```
    * [NgFor](https://angular.io/guide/template-syntax#ngfor)
    * [NgModel - Two-way binding to form elements with [(ngModel)]](https://angular.io/guide/template-syntax#ngModel)
2. 增加過濾條件
    ```ts
    export class AppComponent {
      filterCondition: TodoStatus | undefined;

      // setting
      showAll() { this.filterCondition = undefined; }
      showActive() { this.filterCondition = TodoStatus.Active; }
      showCompleted() { this.filterCondition = TodoStatus.Completed; }

      // check
      get isShowAll() { return !this.filterCondition; }
      get isShowActive() { return this.filterCondition === TodoStatus.Active; }
      get isShowCompleted() { return this.filterCondition === TodoStatus.Completed; }
    }
    ```
    ```html
    <footer class="footer">
      <ul class="filters">
        <li> <a [class.selected]="isShowAll" (click)="showAll()">All</a> </li>
        <li> <a [class.selected]="isShowActive" (click)="showActive()">Active</a> </li>
        <li> <a [class.selected]="isShowCompleted" (click)="showCompleted()">Completed</a> </li>
      </ul>
    </footer>
    ```
    * [Attribute, class, and style bindings](https://angular.io/guide/template-syntax#attribute-class-and-style-bindings)
3. 增加過濾功能
    * 使用 `Function`
        ```ts
        export class AppComponent {
          todoFilterList () {
            if (!this.filterCondition) { return this.todoList; }
            return this.todoList.filter(todo => todo.stauts === this.filterCondition);
          }
        }
        ```
    * 使用 `Pipe`
        ```ts
        @Pipe({ name: 'todoStatusFilter' })
        export class TodoStatusPipe implements PipeTransform {
          transform(todoList: Todo[], status?: TodoStatus): any {
            if (!status || !todoList) { return todoList; }
            return todoList.filter(todo => todo.status === status);
          }
        }
        ```
        * 手動建立
          * 需於 `app.module.ts` 內的 `declarations` 增加對應 `class`
        * CLI: `ng generate pipe [PATH/]NAME`


### Update

1. 連點文字開啟編輯
    ```html
    <li [class.completed]="todo.isCompleted" [class.editing]="todo.selected">
      <div class="view">
        <label (dblclick)="select(todo, edit)">{{todo.name}}</label>
      </div>
      <input class="edit" #edit (keyup.enter)="rename(todo, newName)" (blur)="todo.selected = false" [(ngModel)]="newName" autofocus>
    </li>
    ```
    ```ts
    export class Appcomponent {
      newName: string;

      select(todo: Todo, input: HTMLInputElement) {
        if (todo.isCompleted) { return; }
        todo.selected = true;
        this.newName = todo.name;
        setTimeout(() => input.focus(), 0);
      }

      rename(todo: Todo, newName: string) {
        todo.name = newName;
        todo.selected = false;
      }
    }
    ```

<!--
Q: Template reference variables in ngfor
  [Template reference variable warning notes](https://angular.io/guide/template-syntax#template-reference-variable-warning-notes)
  `The scope of a reference variable is the entire template. Do not define the same variable name more than once in the same template. The runtime value will be unpredictable.`
  REF: https://github.com/angular/angular/issues/13974
-->

### Delete

1. 點 `x` 刪除單筆
    ```html
    <ul class="todo-list" *ngFor="let todo ...; let index = index;">
      <li class="...">
        <div class="view">
          ...
          <button class="destroy" (click)="remove(index)"></button>
        </div>
      </li>
    </ul>
    ```
    * [*ngFor with index](https://angular.io/guide/template-syntax#ngfor-with-index)
    ```ts
    remove(removeIndex: number) {
      this.todoList.splice(removeIndex, 1);
    }
    ```
2. 點 `Clear completed` 刪除多筆
    ```ts
    clearCompleted() {
      this.todoList = this.todoList.filter(todo => !todo.isCompleted);
    }
    ```

## Other

1. 搬移關聯的 `styles` 到 `app.component.css`
2. 文字不要被選取
    ```css
    .todo-list li label {
      ...
      user-select: none;
    }
    ```
