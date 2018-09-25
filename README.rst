a foolish attempt to brute-force a puzzle.

facts:

1. One board with a 3x3 grid of spaces
2. around the edge of the board, 24 tunnels, 2 per side of each empty space
3. each edge-tunnel has a corresponding edge-tunnel elsewhere on the board
4. The "solution" is the one that connects all the corresponding edge-tunnels to each other
5. Nine square tiles, the same size as the empty spaces on the board
6. each tile has two tunnel entries on each of its sides
7. the board has a shadow of a tile in the center space [1,1]

envelope math:

Assuming the center tile is fixed, there are 8 positions to choose.

This gives us ``8!`` possible tile orders, or 40,320.

Each tile can be rotated to any of 4 rotations. A single tile order can have ``4^8`` rotation orders, or 65,536.

``40,320 * 65,536 = 2,642,411,520``

I ran a test of one version of the code, picking random tile orders and random rotation orders (without the center tile fixed), and the program did about 206,000,000 attempts in 7 hours. At that rate, I could test all 2.6 billion permutations in roughly 90 hours.
