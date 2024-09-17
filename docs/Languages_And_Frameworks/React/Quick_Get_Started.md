# React: Quick Get Started

## What is React

It's a library for web a native user interfaces.

You design simple views for each state in your app, and React takes care of
updating and rendering just the right components when your data changes.

Declarative views make your code more predictable and easier to debug.
Declarative is a programming paradigm where you describe *what* your program should do.
SQL, HTML, functional programming languages, and React are examples of declarative programming.

React is also Component-Based. That means you can build encapsulated components that manage
their own state, then compose them to make complex UIs.

React can also render on the server using Node and power mobile apps using React Native.

What are the features of React?

* JSX (JS Syntax Extension)
* TSX (TS Syntax Extension)
* virtual DOM
* the composition of components
* one-way data biding

What is a higher-order components (HOC)?

A higher-order component is a function that takes a component and returns a new
component. It helps reuse component logic.

## Stuff I must know for the interview

* HOOKS:
  * useRef
    * allows you to create a reference to a DOM element or a mutable value that persists across renders
    without triggering re-renders.
    * Use it when you need to directly access or manipulate DOM elements or store a value that does
    not cause re-rendering when updated.
  * useMemo | userCallback | useLayoutEffect | useState | useEffect | useReducer | useContext
  * When to use them, when not to use them
  * what does the second optional argument for useEffect do?
  * What's the difference between useState and useRef?
  * When would you use useState compared to useReducer?
  * What's the difference between useMemo and useCallback?
* STATE MANAGEMENT:
* CONTEXT: useContext
* DATA FETCHING LIBRARIES: rtk
* TSX:
* CUSTOM HOOKS:
  * Why do you need them and how would you write them?
* use effect
* dependencies
* cleanup functions
* How do you build a component?
* What's the approach for extension or changing features?
* Benefits of use reducer vs. useState?
* State modeling
* Why booleans can be bad?
* Architecture. What does architecture mean to you? What is the best approach for architecture in a React project?
* Benefits and pitfalls of css in js.
* What do you think of tailwind? And Vite? and other frameworks?
* What is Rendering cycle?
* Tell me about class vs functional components.
* What's reconciliation in react?
* What do you need to think about when building a modal or according?
* How do you split things up?
* Accessibility concerns to think about (ARIA, semantic html)
* What headless cms's you have used and approaches to this
* Scalability - how do you ensure apps can scale?
* Performance
  * What can you do to improve the performance of a react application?
* Virtual DOM
  * Get to know what this is
  * Get to know how React works in general
  * Memoization
  * lazy loading
  * dynamic imports
* code smells - what do you look for in a PR?
* What new parts of css you have used?
* Benefits and drawbacks of react?
* Learn when it's appropriate to use a hook and when to use an event handler

- TypeScript tutorial - https://www.totaltypescript.com/tutorials

- Why booleans can be bad - https://kyleshevlin.com/enumerate-dont-booleanate

- React architecture - https://profy.dev/article/react-architecture-api-layer

- Using explicit States - https://www.youtube.com/watch?v=ul_3ABrpj64

- Frontend/React questions - https://greatfrontend.com/questions/react

- 100+ React questions - https://github.com/sudheerj/reactjs-interview-questions (some of these questions/answers haven't been updated for functional components so just be aware of that)
- https://www.openbrewerydb.org/

Use this interview kit https://www.developerupdates.com/shop/front-end-developer-job-interview-kit

## Refresher

1. Component types
   1. Functional Components: Stateless, typically use React hooks
   2. Class Components: Stateful, with lifecycle method (less common now)
2. Hooks
   1. useState: Manage local component state
   2. useEffect: Side effects, like fetching data or manually interacting with DOM
   3. useContext: For sharing state globally without passing props down
   4. useMemo/useCallback: Optimize performance by memoizing values or functions
3. State & Props
   1. Props: Data passed from parent to child
   2. State: Internal data managed withing a component
4. Component Lifecycle
   1. Key lifecycle methods: `componentDidMount`, `componentDidUpdate`, `componentWillUnmount`
   2. Functional component equivalents: `useEffect` for side effects
5. Handling Forms:
   1. Controllers vs. uncontrolled components
   2. Using `useState` to manage form inputs

TypeScript basics in React

1. Typing functional components:

```tsx
const MyComponent: React.FC<Props> = ({ propName }) => { ... }
```

2. Props and State Typing:

Define types/interfaces for component props and state:

```tsx
interface Props { name: string }
const Component: React.FC<Props> = ({ name }) => { ... }
```

3. Typing Events

```tsx
const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => { ... }
```

4. Generics

Useful for complex types or utility functions:

```tsx
const identity = <T>(arg: T): T => arg;
```

### Tricky questions

1. How does React's reconciliation process work?

React's reconciliation algorithm (the virtual DOM) allows you efficient updates to the actual DOM.
When state changes, React creates a virtual DOM, compares it to the previous one, and updates only
the changed parts.

2. What are closures and how do they impact React hooks?

A closure is a function that remembers the environment in which it was created. This is relevant for
React hooks, especially `useEffect`, because improper use of dependencies can result in state closures,
that is, when state or props inside an effect doesn't update as expected.

3. What problems can arise with `useEffect`?

* infinite loops can happen when you don't properly declare dependencies
* Memory leaks can happen if you forget to clean up side effects (event listeners, subscriptions)

4. What is the difference between controlled and uncontrolled component?

Controlled component: React controls the form input values via state
Uncontrolled component: Form values are managed by the DOM itself via refs

5. How do you handle performance optimization in React?

* Memoization: React.memo, useMemo, useCallback, can prevent unnecessary re-renders
* Code splitting: Using React.lazy() and Suspense to load components only when needed
* Virtualization: Tools like react-window or react-virtualized for rendering large lists.

6. What is the purpose of keys in React?

Keys help React identify which items have changed, are added, or are removed. The key must be unique among siblings.

7. How would you handle an error in a react component?

* Using Error Boundaries, React allows us to catch JavaScript errors anywhere in the component tree and display a fallback UI

8. Explain the concept of React fiber

* Fiber is React's new reconciliation engine, allowing React to split rendering work into chunks and prioritize updates,
providing better control over scheduling and handling animations, gestures, etc.

#### TypeScript Tricky Questions

1. What is the difference between interface and type in TypeScript?

* Interfaces are extendable and more suited for defining object shapes. Types can define unions, intersections, and complex structures like mapped types, and are more flexible in general.
* Difference: Interfaces can be merged, types cannot.

2. How would you type a React event handler?

```tsx
const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => { ... }
```

3. What are utility types in TypeScript?

* Partial<T>: Makes all properties in T optional.
* Pick<T, K>: Picks a set of properties from T
* Omit<T, K>: Removes a set of properties from T

4. What is a discriminated union in TypeScript, and when would you use it?

A discriminated union is a pattern that combines union types and literal types to create type-safe alternatives to conditionals. Useful for Redux-like patterns:

```tsx
type Share = | { kind: 'circle'; radius: number } | { kind: 'square': sideLength: number };

function area(shape: Shape) {
    switch (shape.kind) {
        case 'circle': return Math.PI * shape.radius ** 2;
        case 'square': return shape.sideLength ** 2?
    }
}
```

5. How would you type a custom React Hook?

```tsx
function useFetch<T>(url: string): { data: T | null, error: string | null} {
    // fetch data logic
}
```

6. Explain how TypeScript helps prevent common errors in React development

TypeScript catches issues such as incorrect prop types, improper state management, or
using non-existent methods. It enforces strict types, making the code more predictable and
reducing runtime errors.

## How to create a React app

React has been designed from the start for gradual adoption.

```bash
sudo apt install npm

npx create-react-app basic-react-app
```

If you want to build a new app or a new website fully with React, we recommend picking one of the
React-powered frameworks popular in the community.

### Production-grade React frameworks

* Next.js
  * Next.js' Pages Router is a full-stack React framework. It's versatile and lets you create React apps of any size.
  * npx create-next-app@latest
* Remix
  * Remix is a full-stack React framework with nested routing.
  * npx create-remix
* Gatsby
  * Gatsby is a React framework for fast CMS-backend websites.
  * npx create-gatsby
* Expo
  * Expo is a React framework that lets you create universal Android, iOS, and web apps with truly native UIs.
  * npx create-expo-app

## Recommended Text Editor Features

* Linting
  * ESLint
* Formatting
  * Prettier
  * Run this on every save

## Using TypeScript

TypeScript is a popular way to add type definitions to JavaScript codebases. Out of the box, TypeScript supports
JSX and you can get full React Web support by adding `@types/react` and `@types/react-dom`.

> Note:
> Every file containing JSX must use the `.tsx` file extension. This is a TypeScript-specific extension that
> tells TypeScript that this file contains JSX.

Writing TypeScript with React is very similar to writing JavaScript with React. The key difference when working
with a component is that you can provide types for your component's props. These types can be used for correctness
checking and providing inline documentation in editors.

```tsx
function MyButton({ title }: { title: string }) {
    return (
        <button>{title}</button>
    );
}
```

This inline syntax is the simplest way to provide types for a component, though once you start to have a few fields
to describe it can become unwieldy. Instead, you can use an `interface` or `type` to describe the component's props:

```tsx
interface MyButtonProps {
    title: string;
    disabled: boolean;
}

function MyButton({ title, disabled }: MyButtonProps) {
    return (
        <button disabled={disabled}>{title}</button>
    );
};
```

## Components

Building blocks of React. They can be functional or class-based.
Functional components are preferred in modern React.

```tsx
const MyComponent: React.FC = () => {
    return <h1>Jesus Saves</h1>;
};
```

Class components are less common in modern React.

React apps are made out of components. A component is a piece of the UI (user interface)  that has its own logic
and appearance. A component can be small as a button, or as large as an entire page.

React components are JavaScript/TypeScript functions that return markup.

```jsx
function MyButton() {
    return (
        <button>I'm a button</button>
    );
}
```

Now that you've declared MyButton, you can nest it into another component.

```jsx
export default runction MyApp() {
    return (
        <div>
            <h1>Welcome to my app</h1>
            <MyButton />
        </div>
    );
}
```

Notice that `<MyButton />` starts with a capital letter. That's how you know it's a React component.
React component names must always start with a capital letter, while HTML tags must be lowercase.

### Writing markup with JSX

The markup syntax you've seen above is called JSX. It's optional, but most React projects use JSX for its
convenience.

JSX is stricter than HTML. You have to close tags like `<br />`. Your components also can't return multiple
JSX tags. You have to wrap them into a shared parent, like `<div>...</div>` or an empty `<>...</>` wrapper:

```jsx
function AboutPage() {
    return (
        <>
            <h1>About</h1>
            <p>Hello there, <br />How do you do?</p>
        </>
    );
}
```

### Adding styles

In React, you specify a CSS class with `className`. It works the same way as the HTML `class` attribute.

```jsx
<img className="avatar" />
```

Then you write the CSS rules for it in a separate CSS file:

```css
.avatar {
    border-radius: 50%;
}
```

React does not prescribe how you add CSS files. In the simplest case, you'll add a `<link>` tag to your HTML.

### Displaying Data

JSX lets you put markup into JavaScript. Curly braces let you "escape back" into JavaScript so that you can
embed some variable from your code and display it to the user.

```jsx
return (
    <h1>
        {user.name}
    </h1>
);
```

You can also "escape into JavaScript" from JSX attributes, but you have to use curly braces instead of quotes.
For example, `className="avatar"` passes the `"avatar"` string as the CSS class, but `src={user.imageUrl}` reads
the JavaScript `user.imageUrl` variable value, and then passes that value as the `src` attribute:

```jsx
return (
    <img
        className="avatar"
        src={user.imageUrl}
    />
);
```

You can put more complex expressions inside the JSX curly braces too, for example, string concatenation:

```jsx
<img
    alt={'Photo of ' + user.name}
/>
```

### Conditional rendering

In React, there is no special syntax for writing conditions. Instead, you'll use the same techniques as
you use when writing regular JavaScript code. For example, you can use an `if` statement to conditionally
include JSX:

```jsx
let content;
if (isLoggedIn) {
    content = <AdminPanel />;
} else {
    content = <LoginForm />;
}
return (
    <div>
        {content}
    </div>
);
```

If you prefer more compact code, you ca use the conditional ? operator. Unlike `if`, it works inside JSX:

```jsx
<div>
    {isLoggedIn ? (
        <AdminPanel />
    ) : (
        <LoginForm />
    )}
</div>
```

When you don't need the `else` branch, you can also use a shorter logical `&&` syntax:

```jsx
<div>
    {isLoggedIn && <AdminPanel />}
</div>
```

### Rendering lists

You will rely on JavaScript features like `for` loop and the array `map()` functions to render lists of
components.

For example, let's say you have an array of products:

```jsx
const products = [
    { title: 'Cabbage', id: 1 },
    { title: 'Garlic', id: 2 },
    { title: 'Apple', id: 3 },
];
```

Inside your component, use the `map()` function to transform an array of products into an array of `<li>` items:

```jsx
const listItems = products.map(product =>
    <li key={product.id}>
        {product.title}
    </li>
);
```

Notice how `<li>` has a `key` attribute. For each item in a list, you should pass a string or a number that uniquely
identifies that item among its siblings. Usually, a key should be coming from your data, such as a database ID.
React uses your keys to know what happened if you later insert, delete, or reorder the items.

## Responding to events

You can respond to events by declaring event handler functions inside your components:

```jsx
function MyButton() {
    function handleClick() {
        alert('You clicked me!');
    }

    return (
        <button onClick={handleClick}>
            Click me
        </button>
    );
}
```

Notice how `onClick={handleClick}` has no parentheses at the end! Do not call the event handler function:
you only need to pass it down. React will call your event handler when the user clicks the button.

## State & Props

State: Use `useState` hook in functional components to manage state.

```tsx
const MyComponent: React.FC = () => {
    const [count, setCount] = useState(0);

    return (
        <div>
            <p>{count}</p>
            <button onClick={() => setCount(count + 1)}>Increment</button>
        </div>
    );
};
```

You'll get two things from `useState`: the current state (`count`), and the function that lets you update it
(`setCount`). You can give them any names, but the convention is to write `[something, setSomething]`.

The first time the button is displayed, `count` will be `0` because you passed `0` to `useState()`. When you want
to change state, call `setCount()` and pass the new value to it, clicking this button will increment the counter:

If you render the same component multiple times, each will get its own state.

Props: Data passed from parent to child components.

```tsx
// This is the child component
interface MyProps {
    name: string;
}

const MyComponent: React.FC<MyProps> = ({ name }) => (
    <div>Hello, {name}</div>
);

export default MyComponent;

// This is the parent component
const Home: React.FC = () => {
    const name = 'John';

    return (
        <div>
            <MyComponent name={name}></User>
        </div>
    );
};

export default Home;
```

### Using Hooks

Functions starting with `use` are called Hooks. `useState` is a built-in Hook provided by React.

You can only call Hooks at the top of your components. If you want to use `useHook` in a condition
or loop, extract a new component and call it there.

### Sharing data between components

To make two sibling components display the same state and update together, you need to move the
state from the individual components "upwards" to the closest component containing all of them.

```jsx
export default function MyApp() {
    const [count, setCount] = useState(0);

    function handleClick() {
        setCount(count + 1);
    }

    return (
        <div>
            <h1>Counters that update together</h1>
            <MyButton count={count} onClick={handleClick} />
            <MyButton count={count} onClick={handleClick} />
        </div>
    );
}
```

The information you pass down like this is called props.

Finally, change `MyButton` to read the props you have passed from its parent component:

```jsx
function MyButton({ count, onClick }) {
    return (
        <button onClick={onClick}>
            Clicked {count} times
        </button>
    );
}
```

This is called lifting state up.
By moving up state, you've shared it between components.

## Notes about the Tutorial

```jsx
export default function Square() {
    return <button className="square">X</button>;
}
```

The first line declares a function called `Square`. The `export` JavaScript keyword makes
this function accessible outside of this file. The `default` keyword tells other files
using your code that it's the main function in your file.

The second line returns a button. The `return` JavaScript keyword means whatever comes after is returned
as a value to the caller of the function. `<button>` is a JSX element. A JSX element is a combination of
JavaScript code and HTML tags that describe what you'd like to display. `className="square"` is a button
property or prop that tells CSS how to style the button. `X` is the test displayed inside of the button
and `</button>` closes the JSX element to indicate that any following content shouldn't be placed inside
the button.

**index.js**

```js
import React, { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./styles.css";

import App from "./App";

const root = createRoot(document.getElementById("root"));
root.render(
  <StrictMode>
    <App />
  </StrictMode>
);
```

Lines 1-5 bring all the necessary pieces together"

* React
* React's library to talk to web browsers (React DOM)
* the styles for your components
* the component you created in `App.js`

The remainder of the file brings all the pieces together and injects the final product into `index.html` in
the `public` folder.

Back to App.js

React components need to return a single JSX elements and not multiple adjacent JSX elements like two
buttons. For that, you can use Fragments (`<>` and `</>`) to wrap multiple adjacent JSX elements like
this:

```js
export default function Square() {
  return (
    <>
      <button className="square">X</button>
      <button className="square">X</button>
    </>
  );
}
```

## Thinking in React

To implement a UI in React, you will usually follow the same five steps.

### Step 1: Break the UI into a component hierarchy

Start by drawing boxes around every component and sub-component in the mockup and naming them.

You can think about splitting up a design into components like a programmer:

> Use the same techniques for deciding if you should create a new function or object.
> One such technique is the single responsibility principle, that is, a component should ideally only do one
> thing. If it ends up growing, is should be decomposed into smaller sub-components.

### Step 2: Build a static version in React

### Step 3: Find the minimal but complete representation of UI state

Think of state as the minimal set of changing data that your app needs to remember.
The most important principle for structuring state is to keep it DRY (Don't Repeat Yourself).
Figure out the absolute minimal representation of the sate your application needs and compute
everything else on-demand. For example, if you're building a shopping list, you can store the
items as an array in state. If you want to also display the number of items in the list, don't
store the number of items as another state value - instead, read the length of your array.

How to identify which pieces of data are state?

* Does it remain unchanged over time? If so, it isn't state.
* Is it passed in from a parent via props? If so, it isn't state.
* Can you compute it based on existing state or props in your component? If so, it definitely isn't state!

What's left is probably state.

> Props vs State
> There are two types of "model" data in React: props and state. The two are very different:
>
> * Props are like arguments you pass to a function. They let a parent component pass data to a child
> component and customize its appearance. For example, a `Form` can pass a `color` prop to a `Button`.
> * State is like a component's memory. It lets a component keep track of some information and change
> and change it in response to interactions. For example, a `Button` might keep track of `isHovered`
> state.
>
> Props and state are different but they work together. A parent component will often keep some
> information in state (so that it can change it), and pass it down to a child component as their
> props.

### Step 4: Identify where your state should live

For each piece of state in your application:

1. Identify every component that renders something based on that state
2. Find their closest common parent component - a component above them all in the hierarchy.
3. Decide where the state should live:
   1. Often, you can put the state directly into their common parent
   2. You can also put the state into some component above their common parent
   3. If you can't find a component where it makes sense to own the state, create a new component
      solely for holding the state and add it somewhere in the hierarchy above the common parent component.


> Hooks are special functions that let you "hook into" React.

### Step 5: Add inverse data flow

Pass functions down from parent to child components.

## Virtual DOM

Virtual DOM is programming concept where an ideal, or "virtual", representation of a UI is kept in memory and synced with the "real" DOM
by a library such as ReactDOM. This process is called reconciliation.