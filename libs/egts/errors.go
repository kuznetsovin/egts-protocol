package egts

//код сообщения, что пакет успешно обработано
const egtsPcOk = uint8(0)

//код сообщения, что пакет в процессе обработки (результат обработки ещё не известен)
const egtsPcInProgress = uint8(1)

//неподдерживаемый протокол
const egtsPcUnsProtocol = uint8(128)

//ошибка декодирования
const egtsPcDecryptError = uint8(129)

//обработка запрещена
const egtsPcProcDenied = uint8(130)

//неверный формат заголовка
const egtsPcIncHeaderform = uint8(131)

//неверный формат данных
const egtsPcIncDataform = uint8(132)

//EgtsPcUnsType неподдерживаемый тип
const egtsPcUnsType = uint8(133)

//неверное количество параметров
const egtsPcNotenParams = uint8(134)

//попытка повторной обработки
const egtsPcDblProc = uint8(135)

//обработка данных от источника запрещена
const egtsPcProcSrcDenied = uint8(136)

//ошибка контрольной суммы заголовка
const egtsPcHeaderCrcError = uint8(137)

//ошибка контрольной суммы данных
const egtsPcDatacrcError = uint8(138)

//некорректная длина данных
const egtsPcInvdatalen = uint8(139)

//маршрут не найден
const egtsPcRouteNfound = uint8(140)

//маршрут закрыт
const egtsPcRouteClosed = uint8(141)

//маршрутизация запрещена
const egtsPcRouteDenied = uint8(142)

//неверный адрес
const egtsPcInvaddr = uint8(143)

//превышено количество ретрансляции данных
const egtsPcTtlexpired = uint8(144)

//нет подтверждения
const egtsPcNoAck = uint8(145)

//объект не найден
const egtsPcObjNfound = uint8(146)

//событие не найдено
const egtsPcEvntNfound = uint8(147)

//сервис не найден
const egtsPcSrvcNfound = uint8(148)

//сервис запрещён
const egtsPcSrvcDenied = uint8(149)

//неизвестный тип сервиса
const egtsPcSrvcUnkn = uint8(150)

//авторизация запрещена
const egtsPcAuthPenied = uint8(151)

//объект уже существует
const egtsPcAlreadyExists = uint8(152)

//идентификатор не найден
const egtsPcIDNfound = uint8(153)

//неправильная дата и время
const egtsPcIncDatetime = uint8(154)

//ошибка ввода/вывода
const egtsPcIoError = uint8(155)

//недостаточно ресурсов
const egtsPcNoResAvail = uint8(156)

//внутренний сбой модуля
const egtsPcModuleFault = uint8(157)

//сбой в работе цепи питания модуля
const egtsPcModulePwrFlt = uint8(158)

//сбой в работе микроконтроллера модуля
const egtsPcModuleProcFlt = uint8(159)

//сбой в работе программы модуля
const egtsPcModuleSwFlt = uint8(160)

//сбой в работе внутреннего ПО модуля
const egtsPcModuleFwFlt = uint8(161)

//сбой в работе блока ввода/вывода модуля
const egtsPcModuleIoFlt = uint8(162)

//сбой в работе внутренней памяти модуля
const egtsPcModuleMemFlt = uint8(163)

//тест не пройден
const egtsPcTestFailed = uint8(164)
