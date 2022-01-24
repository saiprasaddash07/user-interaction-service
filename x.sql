SELECT
    contentId
FROM
    (
        SELECT
            c.contentId AS contentId,
            COUNT(l.likeId) AS likes,
            COUNT(rI.readId) AS readsContent
        FROM
            content c
            JOIN likes l ON l.contentId = c.contentId
            JOIN readInteraction rI ON c.contentId = rI.contentId
    ) AS interaction
ORDER BY
    likes DESC,
    readsContent DESC;