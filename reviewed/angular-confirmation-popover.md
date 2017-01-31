<!-- TOC -->

- [How To Use](#how-to-use)
- [Learned](#learned)
  - [Set Defaults to Other Module](#set-defaults-to-other-module)
  - [Positioning](#positioning)
  - [Renderer Service](#renderer-service)
  - [ViewContainerRef](#viewcontainerref)
    - [Get ViewContainerRef with dependency injection](#get-viewcontainerref-with-dependency-injection)
    - [Get ViewContainerRef with ViewChild](#get-viewcontainerref-with-viewchild)
    - [Create Component with Injector](#create-component-with-injector)
  - [Inject Dcoument](#inject-dcoument)
  - [Host Listen Document Event](#host-listen-document-event)
  - [Host Listen Window Event](#host-listen-window-event)
  - [Template](#template)

<!-- /TOC -->

# How To Use

* Project: [link](https://github.com/mattlewis92/angular-confirmation-popover)
* Docs: [link](https://mattlewis92.github.io/angular-confirmation-popover/docs/)
* Demo: [link](https://mattlewis92.github.io/angular-confirmation-popover/)

```HTML
<button
 class="btn btn-default"
 mwlConfirmationPopover
 [title]="title"
 [message]="message"
 placement="left"
 (confirm)="confirmClicked = true"
 (cancel)="cancelClicked = true"
 [(isOpen)]="isOpen">
  Show confirm popover!
</button>
```


# Learned

## Set Defaults to Other Module

```typescript
// Self Module
@NgModule({
  declarations: [SelfComponent],
  imports: [
    OtherModule.forRoot({ /* set defaults here*/ })
  ],
  bootstrap: [SelfComponent]
})
class SelfModule {}
```


```typescript
// Other Module
export const USER_OPTIONS: OpaqueToken = new OpaqueToken('USER_OPTIONS');

export function optionsFactory(userOptions: UserOptions): UserOptions {
  const options: UserOptions = new UserOptions();
  Object.assign(options, userOptions);
  return options;
}

@NgModule({...})
export class OtherModule {
  static forRoot(options: UserOptions = {}): ModuleWithProviders {
    return {
      ngModule: OtherModule,
      providers: [
        { provide: USER_OPTIONS, useValue: options },
        { provide: UserOptions, useFactory: optionsFactory, deps: [USER_OPTIONS] },
      ]
    };
  }
}
```

## Positioning

* [ng-bootstrap Positioning class as a standalone module](https://github.com/mattlewis92/positioning/)

```typescript
// Copy from README.md
import { Positioning } from 'positioning';

const positioning = new Positioning();
const position = positioning.positionElements(host, target, placment, appendToBody);
```

## Renderer Service

**EXPERIMENTAL** tag in [official docs](https://angular.io/docs/ts/latest/api/core/index/Renderer-class.html).

* [Angular 2 — Explore The Renderer Service](https://netbasal.com/angular-2-explore-the-renderer-service-e43ef673b26c#.3c5ectly9)
* [Custom Renderer for Angular2](https://github.com/ralfstx/angular2-renderer-example)

```typescript
import { Directive, Renderer, ElementRef } from '@angular/core';
@Directive({
  selector: '[exploreRenderer]'
})
export class ExploreRendererDirective {
  private nativeElement : Node;
  constructor( private renderer : Renderer, private element : ElementRef ) {
    this.nativeElement = element.nativeElement;
  }
}
```

```typescript
// Example
this.renderer.invokeElementMethod(this.elm.nativeElement, 'focus', []);
this.renderer.setElementStyle(popoverElement, 'top', '100px');
```

## ViewContainerRef

* [ViewContainerRef official docs](https://angular.io/docs/ts/latest/api/core/index/ViewContainerRef-class.html)
* [Understanding ViewContainerRef in Angular 2](https://netbasal.com/angular-2-understanding-viewcontainerref-acc183f3b682#.6s41s2ozp)

### Get ViewContainerRef with dependency injection

```typescript
@Component({
  selector: 'vcr',
  template: `
    <template #tpl>
      <h1>ViewContainerRef</h1>
    </template>
  `,
})
export class VcrComponent {
  @ViewChild('tpl') tpl;
  constructor(private _vcr: ViewContainerRef) { }

  ngAfterViewInit() { this._vcr.createEmbeddedView(this.tpl); }
}
```

### Get ViewContainerRef with ViewChild

```typescript
@Component({
  selector: 'vcr',
  template: `
    <template #tpl>
      <h1>ViewContainerRef</h1>
    </template>
    <div>Some element</div>
    <div #container></div>
  `,
})
export class VcrComponent {
  @ViewChild('container', { read: ViewContainerRef }) _vcr;
  @ViewChild('tpl') tpl;

  ngAfterViewInit() { this._vcr.createEmbeddedView(this.tpl); }
}
```

### Create Component with Injector

The component is instantiated using its `ComponentFactory` which can be obtained via `ComponentFactoryResolver`.

```typescript
import { ComponentFactoryResolver, ViewContainerRef, ComponentRef, ReflectiveInjector, ResolvedReflectiveProvider, Injector } from '@angular/core';

// use example
const componentFactory: ComponentFactory<SomeComponent> = this.componentFactoryResolver.resolveComponentFactory(SomeComponent);
const binding: ResolvedReflectiveProvider[] = ReflectiveInjector.resolve([{ provide: SomeProvide, useValue: provideValue }]);
const contextInjector: Injector = this.viewContainerRef.parentInjector;
const childInjector: Injector = ReflectiveInjector.fromResolvedProviders(binding, contextInjector);
const childComponent: ComponentRef<SomeComponent> = this.viewContainerRef.createComponent(componentFactory, this.viewContainerRef.length, childInjector);
```

* [ReflectiveInjector official docs](https://angular.io/docs/ts/latest/api/core/index/ReflectiveInjector-class.html)
* [ComponentRef official docs](https://angular.io/docs/ts/latest/api/core/index/ComponentRef-class.html)
* [Injector official docs](https://angular.io/docs/ts/latest/api/core/index/Injector-class.html)

## Inject Dcoument

```typescript
import { DOCUMENT } from '@angular/platform-browser';

@Component({...})
export class SomeComponent {
  constructor(@Inject(DOCUMENT) private document) { }
}
```

## Host Listen Document Event

```typescript
@Directive({...})
export class SomeDirective {
  @HostListener('document:click', ['$event.target'])
  @HostListener('document:touchend', ['$event.target'])
  onDocumentClick(target: HTMLElement): void { }
}
```

## Host Listen Window Event

```typescript
@Directive({...})
export class SomeDirective {
  @HostListener('window:resize')
  onResize(): void { }
}
```

## Template

* [Template official docs](https://angular.io/docs/ts/latest/guide/template-syntax.html)
* [NgTemplateOutlet official docs](https://angular.io/docs/ts/latest/api/common/index/NgTemplateOutlet-directive.html)
* [Template Tag](http://blog.kevinyang.net/2016/11/12/ng2-template/)
* [Angular 2 ng-container Stackoverflow](http://stackoverflow.com/questions/39547858/angular-2-ng-container)

```typescript
// Copy from official doc
@Component({
  selector: 'ng-template-outlet-example',
  template: `
    <ng-container *ngTemplateOutlet="greet"></ng-container>
    <hr>
    <ng-container *ngTemplateOutlet="eng; context: myContext"></ng-container>
    <hr>
    <ng-container *ngTemplateOutlet="svk; context: myContext"></ng-container>
    <hr>
    <template #greet><span>Hello</span></template>
    <template #eng let-name><span>Hello {{name}}!</span></template>
    <template #svk let-person="localSk"><span>Ahoj {{person}}!</span></template>
`
})
class NgTemplateOutletExample {
  myContext = {$implicit: 'World', localSk: 'Svet'};
}
```
