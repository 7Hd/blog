<!-- TOC -->

- [How To Use](#how-to-use)
- [Learned](#learned)
  - [Set Defaults to Other Module](#set-defaults-to-other-module)
  - [Positioning](#positioning)
  - [Renderer Service](#renderer-service)

<!-- /TOC -->

# How To Use

* Project: [link](https://github.com/mattlewis92/angular-confirmation-popover)
* Docs: [link](https://mattlewis92.github.io/angular-confirmation-popover/docs/)
* Demo: [link](https://mattlewis92.github.io/angular-confirmation-popover/)


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
