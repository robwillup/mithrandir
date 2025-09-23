# Blazor course from Pluralsight

## ASP.NET Core:
* wwwroot: this is the location of the static files
* componets: In Blazor, "components" are really import. Anything in Blazor is a component.
* razor: .razor files are components
* To reuse a component in another component, that common SPA behavior, just use the component name surrounded by <>, like:

Suppose you created a component with this name: Counter.razor


```Razor
<Counter />
```

## Using Code

Many consider the "code behind" approach better. So instead of using the @code{} inside the .razor file, create another class using "partial" just for the code.

```CSharp
public partial class EmployeeOverview
{

}
```

## Lifecycle methods
In Blazor if you want to initialize some variables or properties a constructor is not a good idea, instead of should use  the Lifecycle methods, such as `OnInitializedAsync`.