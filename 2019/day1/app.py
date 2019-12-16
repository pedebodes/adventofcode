
# Efetuando a leitura do arquivo TXT e transformando em dictionary
def get_pair(line):
    key, sep, value = line.strip().partition(" ")
    return int(key), value
with open("dados.txt") as fd:    
    dados = dict(get_pair(line) for line in fd)


def somandoAnegada(array):
    import numpy as np
    return sum(map(np.array, array))

def calcula(valor):
    return (valor /3)-2
    # return round((valor /3)-2,0)

# Efetuando Processamento
def fase_1(dados):
    # dados = {135568}
    fase1  = []
    for d in dados:
        fase1.append(calcula(d))
    return fase1

#segunda fase desafio , efetuar o calculo ate zerar
def fase_2(dados):
    # dados = {100756}
    somatorio = 0
    for d in dados:
        
        dados = calcula(d)
        print dados
        print ".."
        somatorio = dados
        while(dados > 0):
            dados = calcula(dados)
            print dados
            print ".."
            if dados > 0:
                somatorio += dados

    return somatorio




fase1 = fase_1(dados)
# print("Resultado da fase 1 e: ",  somandoAnegada(fase1))

print fase_2(fase1)

