# Building Simplicial Homology

Given a **simplicial complex** $K$, simplicial homology is obtained by building chain groups and boundary operators, and then computing quotients of cycles by boundaries.

The procedure goes as follows.

### 1. Identify the simplices by dimension

Split the simplices of $K$ into:

$$
K_0=\{\text{vertices}\},\qquad
K_1=\{\text{edges}\},\qquad
K_2=\{\text{triangles}\},\ldots
$$

For example, suppose:

$$
K_0=\{v_0,v_1,v_2\},
$$

$$
K_1=\{[v_0,v_1],[v_1,v_2],[v_0,v_2]\},
$$

and

$$
K_2=\{[v_0,v_1,v_2]\}.
$$

In this case we have a **filled** triangle.

---

### 2. Build the chain groups $C_k(K)$

For each dimension $k$, we define

$$
C_k(K;\mathbb F)
$$

as the vector space spanned by the $k$-simplices.

If we work over a field $\mathbb F$, for example $\mathbb F_2$,

$$
C_k(K;\mathbb F_2)
=
\operatorname{span}_{\mathbb F_2}(K_k).
$$

In the example:

$$
C_0 \cong \mathbb F_2^3,
$$

$$
C_1 \cong \mathbb F_2^3,
$$

$$
C_2 \cong \mathbb F_2.
$$

A $1$-dimensional chain, for example, has the form

$$
c=a_{01}[v_0,v_1]+a_{12}[v_1,v_2]+a_{02}[v_0,v_2].
$$

---

### 3. Define the boundary operators

For each $k$, we define

$$
\partial_k:C_k\longrightarrow C_{k-1}.
$$

Over $\mathbb Z$, the boundary of an oriented $k$-simplex is

$$
\boxed{
\partial_k[v_0,\ldots,v_k]
=
\sum_{i=0}^{k}
(-1)^i
[v_0,\ldots,\widehat{v_i},\ldots,v_k]
}
$$

where $\widehat{v_i}$ means that $v_i$ is omitted.

For example:

$$
\partial_1[v_0,v_1]
=
[v_1]-[v_0].
$$

And

$$
\partial_2[v_0,v_1,v_2]
=
[v_1,v_2]
-[v_0,v_2]
+[v_0,v_1].
$$

If we use $\mathbb F_2$, the signs disappear because

$$
-1=1\pmod 2.
$$

So:

$$
\partial_2[v_0,v_1,v_2]
=
[v_1,v_2]+[v_0,v_2]+[v_0,v_1].
$$

---

### 4. Assemble the boundary matrices

Once bases are chosen for $C_k$ and $C_{k-1}$, each $\partial_k$ becomes a matrix.

For example, for

$$
C_1=
\langle
e_{01},e_{12},e_{02}
\rangle
$$

and

$$
C_0=
\langle
v_0,v_1,v_2
\rangle,
$$

over $\mathbb F_2$,

$$
\partial_1 e_{01}=v_0+v_1,
$$

$$
\partial_1 e_{12}=v_1+v_2,
$$

$$
\partial_1 e_{02}=v_0+v_2.
$$

Therefore,

$$
[\partial_1]
=
\begin{pmatrix}
1&0&1\\
1&1&0\\
0&1&1
\end{pmatrix}.
$$

For $\partial_2$,

$$
\partial_2[v_0,v_1,v_2]
=
e_{01}+e_{12}+e_{02},
$$

so

$$
[\partial_2]
=
\begin{pmatrix}
1\\
1\\
1
\end{pmatrix}.
$$

These matrices are the computational core of simplicial homology.

---

### 5. Compute the cycles

The $k$-cycles are the chains with no boundary:

$$
\boxed{
Z_k=\ker \partial_k
}
$$

That is,

$$
c\in Z_k
\iff
\partial_k c=0.
$$

In the example, the chain

$$
e_{01}+e_{12}+e_{02}
$$

is a cycle, since

$$
\partial_1(e_{01}+e_{12}+e_{02})=0.
$$

Geometrically, it represents the outline of the triangle.

---

### 6. Compute the boundaries

The $k$-boundaries are the cycles that are boundaries of $(k+1)$-chains:

$$
\boxed{
B_k=\operatorname{im}\partial_{k+1}.
}
$$

In the example,

$$
\partial_2[v_0,v_1,v_2]
=
e_{01}+e_{12}+e_{02}.
$$

Therefore,

$$
B_1
=
\operatorname{span}
\{
e_{01}+e_{12}+e_{02}
\}.
$$

Note the fundamental property:

$$
\boxed{
\partial_k\circ\partial_{k+1}=0
}
$$

and, consequently,

$$
B_k\subseteq Z_k.
$$

In other words: **every boundary is a cycle**.

---

### 7. Form the homology group

The $k$-dimensional homology is

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

The idea is to treat the cycles that are boundaries as trivial.

For the filled triangle:

$$
Z_1
=
\operatorname{span}
\{e_{01}+e_{12}+e_{02}\}.
$$

But also

$$
B_1
=
\operatorname{span}
\{e_{01}+e_{12}+e_{02}\}.
$$

Thus,

$$
H_1=Z_1/B_1=0.
$$

Although there is a cycle, it is the boundary of a filled triangle, so it does not represent a hole.

---

### 8. Compute the Betti numbers

When working over a field,

$$
\beta_k=\dim H_k.
$$

Since

$$
H_k=\ker\partial_k/\operatorname{im}\partial_{k+1},
$$

we have

$$
\boxed{
\beta_k
=
\dim\ker\partial_k
-
\dim\operatorname{im}\partial_{k+1}.
}
$$

By rank–nullity,

$$
\dim\ker\partial_k
=
\dim C_k-\operatorname{rank}\partial_k,
$$

so

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

where $n_k$ is the number of $k$-simplices.

For the filled triangle:

$$
\beta_0=1,
\qquad
\beta_1=0,
\qquad
\beta_2=0.
$$

Therefore,

$$
H_0\cong \mathbb F,
\qquad
H_1=0,
\qquad
H_2=0.
$$

Topologically: there is **one connected component and no holes**.

---

The chain complex is summarized by

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

and homology is computed, in each dimension, as

$$
\boxed{
H_k
=
\frac{\ker\partial_k}
{\operatorname{im}\partial_{k+1}}.
}
$$

That is essentially the whole algorithm: **enumerate simplices $\rightarrow$ assemble boundary matrices $\rightarrow$ compute kernels and images $\rightarrow$ form the quotient**.
