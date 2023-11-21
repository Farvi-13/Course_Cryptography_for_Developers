# Course_Cryptography_for_Developers
# Написання обгортки для зручного використання бібліотеки, що працює з алгеброю на еліптичних кривих
Мета практичного завдання: отримати досвід з використання готової бібліотечної реалізації перетворень в групі точок еліптичної кривої, а також навчитись встановлювати параметри кривої.

# Завдання:
# 1 Написати обгортки до бібліотечних функцій
Напишіть функції-обгортки для обраної вами бібліотеки, що дадуть вам змогу виконувати основні перетворення з точками ЕК зручним для вас чином. Приклади структури точки і методів для роботи з точками наведені нижче.

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func BasePointGGet() (point ECPoint) {} 		//G-generator receiving
func ECPointGen(x, y *big.Int) (point ECPoint) {}	//ECPoint creation
func IsOnCurveCheck(a ECPoint) (c bool) {} 		//DOES P ∈ CURVE?
func AddECPoints(a, b ECPoint) (c ECPoint) {} 	//P + Q
func DoubleECPoints(a ECPoint) (c ECPoint) {} 	//2P	
func ScalarMult(k big.Int, a ECPoint) (c ECPoint) {}	//k * P
func ECPointToString(point ECPoint) (s string) {} 	//Serialize point
func StringToECPoint(s string) (point ECPoint) {} 	//Deserialize point
func PrintECPoint(point ECPoint) {} 			//Print point

# 2 Перевірити коректність роботи перетворень
З використанням власних функцій-обгорток реалізуйте перевірку коректності роботи бібліотеки та відповідних обгорток. Для цього напишіть обчислення простого рівняння. Приклад рівняння і псевдокоду наведений нижче.

k*(d*G) = d*(k*G)

ECPoint G = BasePointGGet()
big.Int k = SetRandom(256)
big.Int d = SetRandom(256)

H1 = ScalarMult(d, G)
H2 = ScalarMult(k, H1)

H3 = ScalarMult(k, G)
H4 = ScalarMult(d, H3)

bool result = IsEqual(H2, H4)

# Деталі реалізації, інструкцію з запуску та результати
Я реалізував всі функції з запропонованих в 1 завданні та написав тести для перевірки коректності їх роботи.
Так як базова бібліотека в Go, а саме crypto/elliptic вже застаріла та не рекомендується до використання, то я використав бібліотеку github.com/btcsuite/btcd/btcec/v2.
Щоб запустити та перевірити код, просто запустіть файл main.
