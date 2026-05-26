## State and useState Hook

State is how components manage changing data. Unlike props (which are read-only), state can be updated when users interact with your app.

Key ideas:
- State is component-specific data that can change over time
- Use the `useState` hook to add state to functional components
- `useState` returns an array with 2 items: `[currentValue, setValueFunction]`
- When you call the setter function, React re-renders the component with the new value

Example:

```react
import { useState } from 'react';

function Counter() {
  // Declare state: 'count' is the current value, 'setCount' updates it
  const [count, setCount] = useState(0); // 0 is the initial value

  return (
    <div>
      <p>You clicked {count} times</p>
      <button onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  );
}
```

Notice:
- `useState(0)` initializes count to 0
- `setCount(count + 1)` updates the state (triggers re-render)
- `onClick` is an event handler that runs when clicked

> General Tip: In JSX, you must return a single parent element. When returning sibling elements, wrap them in `<div></div>`
