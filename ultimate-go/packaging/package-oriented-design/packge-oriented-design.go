package main

func main() {
	// Scalable apps shoud be splitted into two projects

	// Common/kit repo -> repo of packages to share across many projects
	// |-- CONTRIBUTORS -> list of contributors
	// |-- LICENSE -> license
	// |-- README.md -> info about the project
	// |-- cfg/ -> package
	// |-- log/ -> package
	// |-- etc...

	// Application repo -> application repo
	// |-- cmd/ -> were the binaries are going to be
	// |-- internal/ -> for packages that are going to be exclusively for this project,
	//	   |			golang will block any other project trying to import from this directory.
	//     |-- platform/ -> foundational packages, that one day could be part of the kit repo.
	// |-- vendor/ -> for vendoring 3rd party packages.

	// Main packages inside cmd/ can import packages from anywhere.
	// Packages inside internal/ can import from anywhere besides cmd/.
	// Packages inside platform/ can import only from platform/ or vendor.
}
