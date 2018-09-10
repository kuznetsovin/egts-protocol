package main

//код сообщения, что пакет успешно обработано
const egtsPcOk = 0

//код сообщения, что пакет в процессе обработки (результат обработки ещё не известен)
const egtsPcInProgress = 1

//неподдерживаемый протокол
const egtsPcUnsProtocol = 128

//ошибка декодирования
const egtsPcDecryptError = 129

//обработка запрещена
const egtsPcProcDenied = 130

//неверный формат заголовка
const egtsPcIncHeaderform = 131

//неверный формат данных
const egtsPcIncDataform = 132

//EgtsPcUnsType неподдерживаемый тип
const egtsPcUnsType = 133

//неверное количество параметров
const egtsPcNotenParams = 134

//попытка повторной обработки
const egtsPcDblProc = 135

//обработка данных от источника запрещена
const egtsPcProcSrcDenied = 136

//ошибка контрольной суммы заголовка
const egtsPcHeadercrcError = 137

//ошибка контрольной суммы данных
const egtsPcDatacrcError = 138

//некорректная длина данных
const egtsPcInvdatalen = 139

//маршрут не найден
const egtsPcRouteNfound = 140

//маршрут закрыт
const egtsPcRouteClosed = 141

//маршрутизация запрещена
const egtsPcRouteDenied = 142

//неверный адрес
const egtsPcInvaddr = 143

//превышено количество ретрансляции данных
const egtsPcTtlexpired = 144

//нет подтверждения
const egtsPcNoAck = 145

//объект не найден
const egtsPcObjNfound = 146

//событие не найдено
const egtsPcEvntNfound = 147

//сервис не найден
const egtsPcSrvcNfound = 148

//сервис запрещён
const egtsPcSrvcDenied = 149

//неизвестный тип сервиса
const egtsPcSrvcUnkn = 150

//авторизация запрещена
const egtsPcAuthPenied = 151

//объект уже существует
const egtsPcAlreadyExists = 152

//идентификатор не найден
const egtsPcIDNfound = 153

//неправильная дата и время
const egtsPcIncDatetime = 154

//ошибка ввода/вывода
const egtsPcIoError = 155

//недостаточно ресурсов
const egtsPcNoResAvail = 156

//внутренний сбой модуля
const egtsPcModuleFault = 157

//сбой в работе цепи питания модуля
const egtsPcModulePwrFlt = 158

//сбой в работе микроконтроллера модуля
const egtsPcModuleProcFlt = 159

//сбой в работе программы модуля
const egtsPcModuleSwFlt = 160

//сбой в работе внутреннего ПО модуля
const egtsPcModuleFwFlt = 161

//сбой в работе блока ввода/вывода модуля
const egtsPcModuleIoFlt = 162

//сбой в работе внутренней памяти модуля
const egtsPcModuleMemFlt = 163

//тест не пройден
const egtsPcTestFailed = 164
