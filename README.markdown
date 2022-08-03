# The Motoko Package Manager
Motoko is a set of traditionalist SVR4-compatible package managenment tools, but made in a modernist and cleaner way.  
I'm making this in order to learn Go and some data structure concepts.  

Everything here is a big "TO DO", since I'm still learning the basic concepts of
Go itself and the SVR4 package manager specification.  

## Implementation status

This project currently implements the following tools:  

| Command         | Description
| Status |
|-----------------|-------------------------------------------------------------------------|--------|
| ``pkgproto``(1) | Generate a ``prototype``(4) file for input to the ``pkgmk``(1) command. | W.I.P. |  


The following tools shall be implemented in coming days, weeks, months or, in
the worst case scenario, years.  

<table>
<thead>
  <tr>
    <th>Command</th>
    <th>Description</th>
    <th>Status</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>``pkgmk``(1)</td>
    <td>Create an software package.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``pkgadd``(8)</td>
    <td>Install a software package.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``pkgrm``(8)</td>
    <td>Remove a software package.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``pkgask``(8)</td>
    <td>Store answers to a request script.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``pkgtrans``(1)</td>
    <td>"Translate" packages between datastream formats.<br>According to the
specification, it is used for<br>copying packages into a distribution
medium.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``pkgchk``(8)</td>
    <td>Check if a package is impartial.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``pkginfo``(1)</td>
    <td>Displays information about a software package.<br>This shall be simple,
since it is only based on<br>parsing the ``pkginfo``(5) file from a
package.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``pkgparam``(1)</td>
    <td>Displays package parameters, declared at ``pkginfo``(5).<br>In fact,
it's basically like ``pkginfo``(1), but it displays<br>the file instead of
parsing it for displaying separate values.<br>This shall also be simple.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``installf``(8)</td>
    <td>Installs files into an already installed software package.<br>In theory,
this basically just edits the database and copies<br>the file to the file
system.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td>``removef``(8)</td>
    <td>Removes files from an already installed software package.<br>In theory,
this also just edits the database and deletes<br>the file in the file
system.</td>
    <td>Planning</td>
  </tr>
</tbody>
</table>

## You can read more about it in READMEs below:  
* :brazil: [README.pt.markdown (original)](./docs/README.pt.markdown);
* :us: [README.en.markdown](./docs/README.en.markdown);
* :jp: [README.jp.markdown (TO DO)](./docs/README.jp.markdown);
* :taiwan: [README.zh_TW.markdown (TO DO)](./docs/README.zh_TW.markdown).

## Quick information: file formats
* `.e`: 0 bytes file that means literally "exists", add it to a empty directory
  that need in some way to appear in git's source tree;
* `.markdown`: Markdown files;
* `.go`: Source code for programs written in the Go programming language;
* `.ksh`: ksh file script;
* `.pdf`: PDF document (please, use DjVu);
* `.djvu`: DjVu document.
