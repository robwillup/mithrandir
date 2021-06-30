# sprintf

Ok, so `sprintf` is really not that scary 😌:

```c
char f[3], s[3], t[3], fo[3];
scanf("%s", f);
scanf("%s", s);
scanf("%s", t);
scanf("%s", fo);

char ip[15];
int j;
j = sprintf(ip, "%s.", f);
j += sprintf(ip + j, "%s.",s);
j += sprintf(ip + j, "%s.", t);
j += sprintf(ip + j, "%s", fo);

printf("%s", ip);
    
return 0;
  ```
