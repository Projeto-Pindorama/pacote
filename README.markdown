# The Motoko Package Manager
Motoko is a set of traditionalist SVR4-compatible package managenment tools, but made in a modernist and cleaner way.  
I'm making this in order to learn Go and some data structure concepts.  

Everything here is a big "TO DO", since I'm still learning the basic concepts of
Go itself and the SVR4 package manager specification.  

## You can read more about it in READMEs below:  
* :brazil: [README.pt.markdown (original)](./docs/README.pt.markdown);
* :us: [README.en.markdown](./docs/README.en.markdown);
* :jp: [README.jp.markdown (TO DO)](./docs/README.jp.markdown);
* :taiwan: [README.zh_TW.markdown (TO DO)](./docs/README.zh_TW.markdown).

There is also the official documentation, which is served at
https://silicon.pindorama.dob.jp/motoko and has its source at the
[Silicon Tabula](https://github.com/Projeto-Pindorama/Silicon-Tabula) repository.

## Quick information: file formats
* `.e`: 0 bytes file that means literally "exists", add it to a empty directory
  that need in some way to appear in git's source tree;
* `.markdown`: Markdown files;
* `.go`: Source code for programs written in the Go programming language;
* `.ksh`: ksh file script;
* `.pdf`: PDF document (please, use DjVu);
* `.djvu`: DjVu document.

## Implementation status

This project currently implements the following tools:  

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
    <td><tt>pkgproto</tt>(1)</td>
    <td>Generate a <tt>prototype</tt>(4) file for input to the
<tt>pkgmk</tt>(1) command.</td>
    <td>W.I.P.; Almost finished</td>
  </tr>
</tbody>
</table>

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
    <td><tt>pkgmk</tt>(1)</td>
    <td>Create an software package.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>pkgadd</tt>(8)</td>
    <td>Install a software package.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>pkgrm</tt>(8)</td>
    <td>Remove a software package.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>pkgask</tt>(8)</td>
    <td>Store answers to a request script.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>pkgtrans</tt>(1)</td>
    <td>"Translate" packages between datastream formats.<br>According to the
specification, it is used for<br>copying packages into a distribution
medium.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>pkgchk</tt>(8)</td>
    <td>Check if a package is impartial.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>pkginfo</tt>(1)</td>
    <td>Displays information about a software package.<br>This shall be simple,
since it is only based on<br>parsing the <tt>pkginfo</tt>(5) file from a
package.</td>
    <td>W.I.P.</td>
  </tr>
  <tr>
    <td><tt>pkgparam</tt>(1)</td>
    <td>Displays package parameters, declared at <tt>pkginfo</tt>(5).<br>In fact,
it's basically like <tt>pkginfo</tt>(1), but it displays<br>the file instead of
parsing it for displaying separate values.<br>This shall also be simple.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>installf</tt>(8)</td>
    <td>Installs files into an already installed software package.<br>In theory,
this basically just edits the database and copies<br>the file to the file
system.</td>
    <td>Planning</td>
  </tr>
  <tr>
    <td><tt>removef</tt>(8)</td>
    <td>Removes files from an already installed software package.<br>In theory,
this also just edits the database and deletes<br>the file in the file
system.</td>
    <td>Planning</td>
  </tr>
</tbody>
</table>
