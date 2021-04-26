package utils

const SALE string = "SALE"
const RENTAL string = "RENTAL"

// ZAP RULE CONSTANTS
const ZAP_SALE_MIN_VALUE_BY_METER = 3500.00
const ZAP_SALE_MIN_VALUE = 600000.00
const ZAP_RENTAL_MIN_VALUE = 3500.00

// VIVA REAL RULE CONSTANTES
const VIVAREAL_SALE_MIN_VALUE = 700000.00
const VIVAREAL_RENTAL_MAX_VALUE = 4000.00

// Ele apenas é elegível pro portal ZAP:
//// Quando for aluguel e no mínimo o valor for de R$ 3.500,00.
//// Quando for venda e no mínimo o valor for de R$ 600.000,00.
// Ele apenas é elegível pro portal Viva Real:
//// Quando for aluguel e no máximo o valor for de R$ 4.000,00.
//// Quando for venda e no máximo o valor for de R$ 700.000,00.
