VERSION=6.1.3
EXAMPLES=$(VERSION)/examples

.PHONY: all
all: examples

.PHONY: examples
examples: addNodes deleteNode listNodes listHardwareProfiles listSoftwareProfiles
	
addNodes: $(EXAMPLES)/addNodes/addnodes.go
	./xcompile.sh $(EXAMPLES)/addNodes/addnodes.go ./build/$(EXAMPLES) addNodes

deleteNode: $(EXAMPLES)/deleteNode/deletenode.go
	./xcompile.sh $(EXAMPLES)/deleteNode/deletenode.go ./build/$(EXAMPLES) deleteNode

listNodes: $(EXAMPLES)/listNodes/listnodes.go
	./xcompile.sh $(EXAMPLES)/listNodes/listnodes.go ./build/$(EXAMPLES) listNodes

listHardwareProfiles: $(EXAMPLES)/listHardwareProfiles/listhardwareprofiles.go
	./xcompile.sh $(EXAMPLES)/listHardwareProfiles/listhardwareprofiles.go ./build/$(EXAMPLES) listHardwareProfiles

listSoftwareProfiles: $(EXAMPLES)/listSoftwareProfiles/listsoftwareprofiles.go
	./xcompile.sh $(EXAMPLES)/listSoftwareProfiles/listsoftwareprofiles.go ./build/$(EXAMPLES) listSoftwareProfiles

.SILENT: clean
.PHONY: clean
clean:
	$(RM) -rf ./build
