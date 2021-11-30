--- Dia 18: Interpretação de Muitos Mundos ---
Ao se aproximar de Netuno, um sistema de segurança planetária o detecta e ativa um feixe de trator gigante em Triton ! Você não tem escolha a não ser pousar.

Uma varredura da área local revela apenas uma característica interessante: um enorme cofre subterrâneo. Você gera um mapa dos túneis (sua entrada do quebra-cabeça). Os túneis são muito estreitos para se mover na diagonal.

Apenas uma entrada (marcada @) está presente entre as passagens abertas (marcadas .) e as paredes de pedra ( #), mas você também detecta uma variedade de chaves (mostradas em letras minúsculas) e portas (mostradas em letras maiúsculas). As chaves de uma determinada carta abrem a porta da mesma carta: aabre A, babre Be assim por diante. Você não tem certeza de qual chave é necessária para desativar a viga do trator, portanto precisará coletar todas elas .

Por exemplo, suponha que você tenha o seguinte mapa:

#########
#b.A.@.a#
#########
A partir da entrada ( @), você pode acessar apenas uma porta grande ( A) e uma chave ( a). Mover-se para a porta não ajuda, mas você pode mover as 2etapas para coletar a chave, desbloqueando Ano processo:

#########
#b.....@#
#########
Em seguida, você pode mover as 6etapas para coletar a única outra chave b:

#########
#@......#
#########
Portanto, a coleta de todas as chaves levou um total de 8etapas.

Aqui está um exemplo maior:

########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
A única jogada razoável é pegar a chave ae destrancar a porta A:

########################
#f.D.E.e.C.b.....@.B.c.#
######################.#
#d.....................#
########################
Em seguida, faça o mesmo com a chave b:

########################
#f.D.E.e.C.@.........c.#
######################.#
#d.....................#
########################
... e o mesmo com a chave c:

########################
#f.D.E.e.............@.#
######################.#
#d.....................#
########################
Agora, você pode escolher entre as teclas de e. Embora a chave eesteja mais próxima, coletá-la agora seria mais lenta a longo prazo do que coletar a chave dprimeiro, então essa é a melhor escolha:

########################
#f...E.e...............#
######################.#
#@.....................#
########################
Por fim, colete a chave epara destrancar a porta Ee colete a chave f, executando um total de 86etapas.

Aqui estão mais alguns exemplos:

########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
Caminho mais curto é 132passos: b, a, c, d, f, e,g

#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
Caminhos mais curtos são 136etapas;
um é: a, f, b, j, g, n, h, d, l, o, e, p, c, i, k,m

########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
Caminhos mais curtos são 81etapas; um é: a, c, f, i, d, g, b, e,h

Quantas etapas é o caminho mais curto que coleta todas as chaves?

Sua resposta quebra-cabeça foi 2684.

--- Parte dois ---
Você chega ao cofre apenas para descobrir que não há um cofre, mas quatro - cada um com sua própria entrada.

No seu mapa, encontre a área no meio que se parece com isso:

...
.@.
...
Atualize seu mapa para usar os dados corretos:

@#@
###
@#@
Essa alteração dividirá seu mapa em quatro seções separadas, cada uma com sua própria entrada:

#######       #######
#a.#Cd#       #a.#Cd#
##...##       ##@#@##
##.@.##  -->  #######
##...##       ##@#@##
#cB#Ab#       #cB#Ab#
#######       #######
Como algumas das chaves são para portas em outros cofres, levaria muito tempo para coletar todas as chaves sozinho. Em vez disso, você implanta quatro robôs com controle remoto. Cada um começa em uma das entradas ( @).

Seu objetivo ainda é coletar todas as chaves em menos etapas , mas agora, cada robô tem sua própria posição e pode se mover de forma independente. Você pode controlar remotamente apenas um único robô de cada vez. A coleta de uma chave desbloqueia instantaneamente as portas correspondentes, independentemente do cofre em que a chave ou porta é encontrada.

Por exemplo, no mapa acima, o robô superior esquerdo primeiro coleta a chave a, destrancando a porta Ano cofre inferior direito:

#######
#@.#Cd#
##.#@##
#######
##@#@##
#cB#.b#
#######
Em seguida, o robô inferior direito coleta a chave b, destrancando a porta Bno cofre inferior esquerdo:

#######
#@.#Cd#
##.#@##
#######
##@#.##
#c.#.@#
#######
Em seguida, o robô inferior esquerdo coleta a chave c:

#######
#@.#.d#
##.#@##
#######
##.#.##
#@.#.@#
#######
Finalmente, o robô superior direito coleta a chave d:

#######
#@.#.@#
##.#.##
#######
##.#.##
#@.#.@#
#######
Neste exemplo, foram necessárias apenas algumas 8etapas para coletar todas as chaves.

Às vezes, vários robôs podem ter chaves disponíveis ou um robô pode precisar aguardar a coleta de várias chaves:

###############
#d.ABC.#.....a#
######@#@######
###############
######@#@######
#b.....#.....c#
###############
Em primeiro lugar, no canto superior direito, inferior esquerdo e inferior direito robôs se revezam recolhendo chaves a, be c, um total de 6 + 6 + 6 = 18passos. Então, o robô no canto superior esquerdo pode acessar a tecla d, passando outras 6etapas; a coleta de todas as chaves aqui requer um mínimo de 24etapas.

Aqui está um exemplo mais complexo:

#############
#DcBa.#.GhKl#
#.###@#@#I###
#e#d#####j#k#
###C#@#@###J#
#fEbA.#.FgHi#
#############
O robô superior esquerdo coleta a chave a.
O robô inferior esquerdo coleta a chave b.
O robô superior esquerdo coleta a chave c.
O robô inferior esquerdo coleta a chave d.
O robô superior esquerdo coleta a chave e.
O robô inferior esquerdo coleta a chave f.
O robô inferior direito coleta a chave g.
O robô superior direito coleta a chave h.
O robô inferior direito coleta a chave i.
O robô superior direito coleta a chave j.
O robô inferior direito coleta a chave k.
O robô superior direito coleta a chave l.
No exemplo acima, o menor número de etapas para coletar todas as chaves é 32.

Aqui está um exemplo com mais opções:

#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba@#@BcIJ#
#############
#nK.L@#@G...#
#M###N#H###.#
#o#m..#i#jk.#
#############
Uma solução com o menor número de etapas é:

O robô superior esquerdo coleta a chave e.
O robô superior direito coleta a chave h.
O robô inferior direito coleta a chave i.
O robô superior esquerdo coleta a chave a.
O robô superior esquerdo coleta a chave b.
O robô superior direito coleta a chave c.
O robô superior esquerdo coleta a chave d.
O robô superior esquerdo coleta a chave f.
O robô superior esquerdo coleta a chave g.
O robô inferior direito coleta a chave k.
O robô inferior direito coleta a chave j.
O robô superior direito coleta a chave l.
O robô inferior esquerdo coleta a chave n.
O robô inferior esquerdo coleta a chave m.
O robô inferior esquerdo coleta a chave o.
Este exemplo requer pelo menos 72etapas para coletar todas as chaves.

Após atualizar seu mapa e usar os robôs com controle remoto, quais são as poucas etapas necessárias para coletar todas as chaves?