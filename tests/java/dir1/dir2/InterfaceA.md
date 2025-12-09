```mermaid
classDiagram
	class InterfaceA {
		<<interface>>
		+ void publicVoidMethod()
		+ String publicStringMethod()
		+ int publicIntMethod(String arg)
		+ List~String~ publicGenericMethod(Map~String. Integer~ map)
	}

```