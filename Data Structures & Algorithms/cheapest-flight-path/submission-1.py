import heapq
from typing import List


class Solution:
    def findCheapestPrice(
        self,
        n: int,
        flights: List[List[int]],
        src: int,
        dst: int,
        k: int
    ) -> int:

        # Build adjacency list
        adj = [[] for _ in range(n)]
        for u, v, price in flights:
            adj[u].append((v, price))

        INF = float("inf")

        # dist[city][edgesUsed] = min cost
        # Maximum allowed edges = k + 1
        dist = [[INF] * (k + 2) for _ in range(n)]
        dist[src][0] = 0

        # (cost, city, edgesUsed)
        minHeap = [(0, src, 0)]

        while minHeap:
            cost, node, edgesUsed = heapq.heappop(minHeap)

            # Cheapest valid path found
            if node == dst:
                return cost

            # Ignore stale heap entries
            if cost > dist[node][edgesUsed]:
                continue

            # Cannot take more flights
            if edgesUsed == k + 1:
                continue

            for nei, price in adj[node]:
                nextCost = cost + price
                nextEdges = edgesUsed + 1

                if nextCost < dist[nei][nextEdges]:
                    dist[nei][nextEdges] = nextCost
                    heapq.heappush(
                        minHeap,
                        (nextCost, nei, nextEdges)
                    )

        return -1