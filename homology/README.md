# Construção da Homologia Simplicial

Dado um **complexo simplicial** $K$, a homologia simplicial é obtida construindo grupos de cadeias, operadores bordo e depois calculando quocientes entre ciclos e bordos.

O procedimento é este.

### 1. Identifique os simplexos por dimensão

Separe os simplexos de $K$ em:

$$
K_0=\{\text{vértices}\},\qquad
K_1=\{\text{arestas}\},\qquad
K_2=\{\text{triângulos}\},\ldots
$$

Por exemplo, suponha:

$$
K_0=\{v_0,v_1,v_2\},
$$

$$
K_1=\{[v_0,v_1],[v_1,v_2],[v_0,v_2]\},
$$

e

$$
K_2=\{[v_0,v_1,v_2]\}.
$$

Nesse caso temos um triângulo **preenchido**.

---

### 2. Construa os grupos de cadeias $C_k(K)$

Para cada dimensão $k$, definimos

$$
C_k(K;\mathbb F)
$$

como o espaço vetorial gerado pelos $k$-simplexos.

Se trabalharmos sobre um corpo $\mathbb F$, por exemplo $\mathbb F_2$,

$$
C_k(K;\mathbb F_2)
=
\operatorname{span}_{\mathbb F_2}(K_k).
$$

No exemplo:

$$
C_0 \cong \mathbb F_2^3,
$$

$$
C_1 \cong \mathbb F_2^3,
$$

$$
C_2 \cong \mathbb F_2.
$$

Uma cadeia $1$-dimensional, por exemplo, tem a forma

$$
c=a_{01}[v_0,v_1]+a_{12}[v_1,v_2]+a_{02}[v_0,v_2].
$$

---

### 3. Defina os operadores bordo

Para cada $k$, definimos

$$
\partial_k:C_k\longrightarrow C_{k-1}.
$$

Sobre $\mathbb Z$, o bordo de um $k$-simplexo orientado é

$$
\boxed{
\partial_k[v_0,\ldots,v_k]
=
\sum_{i=0}^{k}
(-1)^i
[v_0,\ldots,\widehat{v_i},\ldots,v_k]
}
$$

onde $\widehat{v_i}$ significa que $v_i$ é removido.

Por exemplo:

$$
\partial_1[v_0,v_1]
=
[v_1]-[v_0].
$$

E

$$
\partial_2[v_0,v_1,v_2]
=
[v_1,v_2]
-[v_0,v_2]
+[v_0,v_1].
$$

Se usarmos $\mathbb F_2$, os sinais desaparecem porque

$$
-1=1\pmod 2.
$$

Então:

$$
\partial_2[v_0,v_1,v_2]
=
[v_1,v_2]+[v_0,v_2]+[v_0,v_1].
$$

---

### 4. Monte as matrizes de bordo

Escolhendo bases para $C_k$ e $C_{k-1}$, cada $\partial_k$ vira uma matriz.

Por exemplo, para

$$
C_1=
\langle
e_{01},e_{12},e_{02}
\rangle
$$

e

$$
C_0=
\langle
v_0,v_1,v_2
\rangle,
$$

sobre $\mathbb F_2$,

$$
\partial_1 e_{01}=v_0+v_1,
$$

$$
\partial_1 e_{12}=v_1+v_2,
$$

$$
\partial_1 e_{02}=v_0+v_2.
$$

Portanto,

$$
[\partial_1]
=
\begin{pmatrix}
1&0&1\\
1&1&0\\
0&1&1
\end{pmatrix}.
$$

Para $\partial_2$,

$$
\partial_2[v_0,v_1,v_2]
=
e_{01}+e_{12}+e_{02},
$$

logo

$$
[\partial_2]
=
\begin{pmatrix}
1\\
1\\
1
\end{pmatrix}.
$$

Essas matrizes são a parte computacional central da homologia simplicial.

---

### 5. Calcule os ciclos

Os $k$-ciclos são cadeias sem bordo:

$$
\boxed{
Z_k=\ker \partial_k
}
$$

Isto é,

$$
c\in Z_k
\iff
\partial_k c=0.
$$

No exemplo, a cadeia

$$
e_{01}+e_{12}+e_{02}
$$

é um ciclo, pois

$$
\partial_1(e_{01}+e_{12}+e_{02})=0.
$$

Geometricamente, ela representa o contorno do triângulo.

---

### 6. Calcule os bordos

Os $k$-bordos são aqueles ciclos que são bordos de simplexos de dimensão $k+1$:

$$
\boxed{
B_k=\operatorname{im}\partial_{k+1}.
}
$$

No exemplo,

$$
\partial_2[v_0,v_1,v_2]
=
e_{01}+e_{12}+e_{02}.
$$

Portanto,

$$
B_1
=
\operatorname{span}
\{
e_{01}+e_{12}+e_{02}
\}.
$$

Observe a propriedade fundamental:

$$
\boxed{
\partial_k\circ\partial_{k+1}=0
}
$$

e, consequentemente,

$$
B_k\subseteq Z_k.
$$

Ou seja: **todo bordo é um ciclo**.

---

### 7. Forme o grupo de homologia

A homologia de dimensão $k$ é

$$
\boxed{
H_k(K;\mathbb F)
=
\frac{Z_k}{B_k}
=
\frac{\ker\partial_k}
{\operatorname{im}\partial_{k+1}}.
}
$$

A ideia é identificar como triviais os ciclos que são bordos.

No triângulo preenchido:

$$
Z_1
=
\operatorname{span}
\{e_{01}+e_{12}+e_{02}\}.
$$

Mas também

$$
B_1
=
\operatorname{span}
\{e_{01}+e_{12}+e_{02}\}.
$$

Assim,

$$
H_1=Z_1/B_1=0.
$$

Embora exista um ciclo, ele é o bordo de um triângulo preenchido e, portanto, não representa um buraco.

---

### 8. Calcule os números de Betti

Quando trabalhamos sobre um corpo,

$$
\beta_k=\dim H_k.
$$

Como

$$
H_k=\ker\partial_k/\operatorname{im}\partial_{k+1},
$$

temos

$$
\boxed{
\beta_k
=
\dim\ker\partial_k
-
\dim\operatorname{im}\partial_{k+1}.
}
$$

Usando rank-nullity,

$$
\dim\ker\partial_k
=
\dim C_k-\operatorname{rank}\partial_k,
$$

logo

$$
\boxed{
\beta_k
=
n_k
-
\operatorname{rank}\partial_k
-
\operatorname{rank}\partial_{k+1}
}
$$

onde $n_k$ é o número de $k$-simplexos.

Para o triângulo preenchido:

$$
\beta_0=1,
\qquad
\beta_1=0,
\qquad
\beta_2=0.
$$

Portanto,

$$
H_0\cong \mathbb F,
\qquad
H_1=0,
\qquad
H_2=0.
$$

Topologicamente: há **uma componente conexa e nenhum buraco**.

---

A cadeia de complexos fica resumida por

$$
\boxed{
\cdots
\longrightarrow
C_2
\xrightarrow{\partial_2}
C_1
\xrightarrow{\partial_1}
C_0
\xrightarrow{\partial_0}
0
}
$$

e a homologia é calculada, em cada dimensão, por

$$
\boxed{
H_k
=
\frac{\ker\partial_k}
{\operatorname{im}\partial_{k+1}}.
}
$$

Esse é essencialmente todo o algoritmo: **enumerar simplexos $\rightarrow$ montar matrizes de bordo $\rightarrow$ calcular núcleos e imagens $\rightarrow$ formar o quociente**.