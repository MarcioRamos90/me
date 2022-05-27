#include <stdio.h>
#include <string.h>

int main() {
	char * firstname = "Jhon";
	char lastname[] = "Doe";

	char name[100];
	
	lastname[0] = 'B';
	
	sprintf(name, "%s %s", firstname, lastname);
	if (strncmp(name, "Jhon Boe", 100) == 0) {
		printf("Done!");
	}
	name[0] = '\0';
	strncat(name, firstname, 2);
	strncat(name, lastname, 3);
	printf("%s\n", name);
	return 0;
}
