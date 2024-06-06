# O Gerenciador de Pacotes... Pacote

## O que é?
 
"Pacote" é um conjunto de ferramentas que permitem que você instale, remova,
crie e distribua pacotes.  O objetivo é oferecer um gerenciador de pacotes que
siga, __de fato__, o modelo de gerenciamento de pacotes SRV4, mas que ao mesmo
tempo não seja resultado de "enxugamento de gelo" de código com mais de 20
anos --- no caso, o código das pkgtools originais da Sun que fora liberado junto
com o OpenSolaris nas quebradas de 2005. Além disso, também queremos
redocumentar o padrão SVR4 para os dias de hoje, além de fazer algo que seja uma
boa referência para quem quiser estudar o funcionamento de um gerenciador de
pacotes.

### Por que a mudança de nome?

Bem, como diria o meu amigo Samuel, é "trivia desnecessária", mas vamos lá: o
nome criava confusão e não era direto acerca do que se tratava o projeto, além
de que não tinha motivo para ter mais uma "referência sociocultural" do que já
tem normalmente no Projeto Pindorama por inteiro, então decidi colocar "pacote"
pois, além de ser uma palavra que só existe nessa forma no português, é direto
sobre o que se trata.

### SVR4? Espera, o que é isso?

"SVR4" é um acrônimo para "System V, revision 4", que foi uma versão do UNIX
lançada em 1987, oriunda de trabalho da AT&T e da Sun Microsystems em conjunto.  
Mesmo sendo considerado o "estopim para a insanidade" do UNIX, pelo menos uma
coisa positiva trouxe que, no caso, foram as melhorias para as pkgtools dos
BSDs<sup>(carece de fontes exatas)</sup>, dando origem ao que se conhece como
"gerenciamento de pacotes SVR4".   

Estou me baseando em diversos documentos (que podem ser encontrados em
[docs/references/](./references)), mas de todos o que mais está sendo consultado
é o "Application Packaging Developer's Guide" da Sun Microsystems, publicado em
Fevereiro de 2000
([805-6338.pdf](https://www.uvm.edu/~fcs/Doc/Solaris8/805-6338.pdf)); em segundo
lugar vêm as manpages do illumos ([illumos.org/man/](https://illumos.org/man)).  

## Objetivos (em uma listinha)

* fidelidade às pkgtools SVR4 originais, inclusive a nível de usabilidade;
* ao mesmo tempo que haja fidelidade, não deve-se haver cumplicidade com erros:
  ou seja, se algo existir nas pkgtools originais mas que possa ser substituído sem
  perdas significativas por algo melhor e mais simples, será;
* cortar dependência o máximo possível com o Shell e com ferramentas no
  ``$PATH``;
* manter a base de código legível e fácil de fazer *hacks*/manutenção;  
* `mk` (plan9port) como sistema de montagem.  

## Não-objetivos (também em uma listinha)

* resolução de dependências;
* _download_ por rede;
* compressão nativa;
* suporte para outras línguas além do inglês em mensagens na tela (isso pode
mudar);  
* ~~[ligações
simbólicas](http://doc.cat-v.org/plan_9/4th_edition/papers/lexnames).~~ Tá,
vamos lá, não estamos no Plan 9. 

Tais coisas devem ficar a cargo de uma abstração.

## Por que não as pkgtools do Slackware ou...?

As pkgtools do Slackware, mesmo sendo incrivelmente boas para algo feito em uma
linguagem para meros hacks e estando numa licença extremamente liberal, ainda
pecam em vários sentidos ao meu ver.  Um deles é também seu calcanhar de Aquiles,
que é sua dependência completa no GNU Bourne-Again Shell, além da ilegibilidade
do código por conta de hacks combinada com a simples falta de semancol em algumas
partes.  

Sobre as pkgtools originais, que a Sun liberou em 2000, simplesmente não vale a
pena usar.  Vários recursos foram cortados no port do Heirloom por serem
"Solaris only", além da base de código ser __realmente__ antiga, depender do
Shell para coisas que poderiam ser feitas direto em C e muito provavelmente
conter __vários__ problemas de segurança.  Ah, eu falei do fato de usarem SQL só
para armazenarem coisas como a lista de pacotes? Pois então.  É algo não só
extremamente antigo, mas também "gordo".
Mesmo que eu não seja exatamente um bom programador, eu acredito que eu vá conseguir
fazer algo mais simples e eficiente usando uma linguagem moderna.  

E sobre outras ferramentas de gerenciamento de pacotes consideradas "kiss", eu
não acho que são sérias, até porque a maioria surgiu em distribuições edgy/meme
e... creio que não precisemos chutar cachorro morto e nem mandar pedrada para a
caixa de correios dos outros.

Há o [BigDL](https://github.com/xplshn/bigdl) de nosso hermano
[xplshn](https://github.com/xplshn), mas são coisas diferentes a partir do
momento que esse projeto foca no uso de sistema em si, enquanto o BigDL,
ignorância à parte minha, funciona de forma mais similar ao Homebrew do macOS,
com binários prontos instalados a partir de um repositório. Seria uma comparação
injusta, pois seria como comparar uma jamanta com um carro esporte: um é utilitário,
o outro visa o conforto e praticidade.

## Quem usa?

No exato momento em que estou escrevendo esse README, ainda ninguém -- afinal,
nem está pronto ainda. :rofl:  Estou criando-o para o Copacabana Linux, mas meu
objetivo é que qualquer um possa usar.  Talvez até mesmo o pessoal do Musl LFS,
que ainda está preso às pkgtools do Slackware levemente modificadas, acabe
utilizando -- e eu ficaria extremamente feliz se o fizessem. :smiley:  

## Padrão de desenvolvimento

Em 2021 estive conversando com meu amigo Vitor Almeida, que manteve o website
do projeto, e chegamos à conclusão de que precisamos manter um padrão no projeto em si.  
Decidi impôr padrões para que não aconteça o que aconteceu constantemente no Otto e em
projetos amadores de quem nunca teve que lidar com VCS, que é um programador "atropelar"
o outro, commits se perderem etc.  
Essa parte possivelmente vai ser bem dividida numa documentação posteriormente, mas
por hora vou deixar no README.  

### Padrão de *commits*

Especificação Conventional Commits 1.0.0:  
https://www.conventionalcommits.org/pt-br/v1.0.0/

### Editores de código usados no desenvolvimento

- acme/acme2k (9fans.github.io/plan9port);
- Vi Improved (vim.org).

## Eu deveria ter criado um FAQ separado?

Possivelmente...  
