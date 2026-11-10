package assertions

import (
    "testing"

    "github.com/stretchr/testify/require"

    "api-tests-template/internal/managers/advertisements/models"
)

func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
    t.Helper()

    for _, item := range items {
        if item.Title == title {
            return
        }
    }

    require.Failf(t, "Объявление не найдено", "В ответе нет объявления с названием %q", title)
}