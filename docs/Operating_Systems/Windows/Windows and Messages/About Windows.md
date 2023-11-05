# About Windows

[Reference](https://learn.microsoft.com/en-us/windows/win32/winmsg/about-windows)

## Desktop Window

When you start the system, it automatically creates the desktop window. The
desktop window is a system-defined window that paints the background of the screen
and serves as the base for all windows displayed by all applications.

The desktop window uses a bitmap to paint the background of the screen. The pattern
created by the bitmap is called the desktop wallpaper. By default, the desktop
window uses the bitmap from a .bmp file specified in the registry as the desktop
wallpaper.

The *GetDesktopWindow* function returns a handle to the desktop window.

A system configuration application, such as Control Panel item, changes the desktop
wallpaper by using the *SystemParametersInfo* function with the *wAction*
parameter set to **SPI_SETDESKWALLPAPER** and the *lpvParam* parameter specifying a
a bitmap file name. *SystemParametersInfo* then loads the bitmap from the specified
file, uses the bitmap to paint the background of the screen, and enters the new
file name in the registry.

The image in the background of the Windows desktop is in a file with the .bmp
format. The name of that file is kept in the Registry. The Windows API provides
functions to change this background image and to get a handle to the desktop
window.

## Application Windows

If your application is a graphical Windows-based application it will create at
least one window, even if it's not visible. That one window is called the
*main window*. This main window servers as the primary interface between the user
and the application.

When an application is started, the system also associates a taskbar button with
the application. That button contains the application icon and title. When the
application is active, its taskbar button is displayed in the pushed state.

## Client Area

The client area is the part of a window where the application displays output,
such as text or graphics.

## Nonclient Area

The title bar, window menu, minimize and maximize buttons, sizing border, and
scroll bars are referred to collectively as the window's nonclient area. The
system manages most aspects of the nonclient area; the application manages the
appearance and behavior of its client area.

## Controls and Dialog Boxes

A control is a window that an application uses to obtain a specific piece of
information from the user, such as the name of a file to open of the desired
point size of a text selection.

Controls are always used in conjunction with another windows - typically, a dialog
box. A dialog box is a window that contains one or more controls. An application
uses a dialog box to prompt the user for input needed to complete a command.
Dialog boxes do not typically use the same set of window components as does a
main window.

A message box is a special box that displays a note, caution, or warning to the
user.

## Window Attributes

An application must provide the following information when creating a window.

### Class Name

You cannot have a window that does not belong to a window class. And an
application that wants to create a window, must first register a window class for
it. Most aspects of a window's appearance and behavior are defined in the window
class. The most important component of a window class is the *window procedure*,
which is a function that receives and processes all input and requests sent to
the window. The system provides the input and requests in the form of messages.

### Window Name

The text string that identifies a window for the user.

### Window Style

Aspects of a window's appearance and behavior that are not specified by the
window's class are defined by the window style which is a named constant.

### Extended Windows Style

Styles not defined by the window class or styles.

### Position

A window's position is defined as the coordinates of its upper left corner.
These coordinates, sometimes called window coordinates, are always relative to the
upper left corner of the screen or, for a child window, the upper left corner
of the parent window's client area. For example, a top-level window having the
coordinates (10,10) is placed 10 pixels to the right of the upper left corner of
the screen and 10 pixels down from it.

The **WindowFromPoint** function returns a handle to the window occupying a
particular point of the screen.

### Size

A window's size (width and height) is given in pixels.

### Parent or Owner Window Handle

A window can have a parent window. A window that has a parent is called a child
window.

A window that has no parent, or whose parent is the desktop window, is called a
top-level window. An application can use the **EnumWindows** function to obtain
a handle to each top-level window on the screen.

A top-level window can own, or be owned by, another window. An owned window always
appears in front of its owner window, is hidden when its owner window is minimized,
and is destroyed when its owner window is destroyed.

### Application Instance Handle

Every application has an instance handle associated with it. The system provides
the instance handle to an application when the application starts. Because it can
run multiple copies of the same application, the system users instance handles
internally to distinguish one instance of an application from another. The
application must specify the instance handle in many different windows, including
those that create windows.

### Creation Data

Every window can have application-defined creation data associated with it. When
the window is first created, the system passes a pointer to the data on to the 
window procedure of the window being created. The window procedure uses the data
to initialize application-defined variables.

### Window Handle

After creating a window, the creation function returns a window handle that
uniquely identifies the window. A window handle has the **HWND** data type; an
application must use this type when declaring a variable that holds a window
handle. An application uses this handle in other functions to direct their actions
to the window.

An application can use the **FindWindow** function to discover whether a window
with the specified class name or window name exists in the system. If such a
window exists, **FindWindow** returns a handle to the window.

## Window-Creation Messages

When creating any window, the system sends messages to the window procedure for
the window.

## Multithreaded Applications

A Windows-based application can have multiple threads of execution, and each
thread can create windows. The thread that creates a window must contain the
code for its window procedure.

An application can use the **EnumThreadWindows** function to enumerate the
windows created by a particular thread.
