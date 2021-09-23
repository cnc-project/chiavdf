package main

// #cgo LDFLAGS: -L . -lcgo -lstdc++ -lm -lgmp -lzcnt
// #cgo CFLAGS: -I ./
// #include "c_interface.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func create_discriminant() bool {
	challenge_hash := C.CString("c86c9946c9d5e221d6e9")
	result := C.c_create_discriminant(challenge_hash, C.int(512))

	defer C.free(unsafe.Pointer(result))
	fmt.Println("discriminant:", C.GoString(result))

	if C.GoString(result) == "-0xff04fc6024a16e52007004e19bb5952b2197b3a3dd1fb72e4866e2c4c5736fa814e9291ed2ff482179c256aa3239bf6869db4eadb2326e48c52f833b3afc9617" {
		return true
	} else {
		return false
	}
}

func prove() bool {
	challenge_hash := "5dcc1975399b860aac6a"
	initial_el := "08000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	discriminant_size_bits := 512
	num_iterations := 100
	result := C.c_prove(C.CString(challenge_hash), C.CString(initial_el), C.int(discriminant_size_bits), C.uint64_t(num_iterations))

	defer C.free(unsafe.Pointer(result))
	fmt.Println("prove:", C.GoString(result))

	if C.GoString(result) == "03006d1a44b489e6ad64be803e28302d3b7683ad89cc679406fa7ccfc1be3160961695d518bd8deb49fe1f13986023dfff30020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" {
		return true
	} else {
		return false
	}
}

func get_b_from_n_wesolowski() bool {
	discriminant := "-0xe7cf6525805e96fffae54fd3887358c7d44b47ad9a58f816b3a0a8b44ce94de25c2f52733ae0bbba8800e03acc1ee7111ef3c22528235a58cc81976de19f789f"
	x := "02003553fe89825c1c93157f5190f24e6b29b86b3c0be0a0543067a0cd1347e3790fcf05ff30e8daa924dd15cb32704ed32f0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	y_result := "0100f427f3e4490f946f7ce0adc34c29da750e380a7dc3b58336ff4e710903602612e33038c6bd1d73a1ea470319b8f152350100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	y_proof := "03002abfcda9a55b09e90f09ce49654ec4a2402d761c25bd3f368aba98bdd9be3005f1567e210f3f61061222280f913a0f100200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	iters := 166670
	result := C.c_get_b_from_n_wesolowski(C.CString(discriminant), C.CString(x), C.CString(y_result+y_proof), C.uint64_t(iters), C.uint64_t(0))

	defer C.free(unsafe.Pointer(result))
	fmt.Println("b:", C.GoString(result))

	if C.GoString(result) == "0xab34f0550be1e0fe666e83bd4596a97d286195c5eada86be0cfbdd69245a7865c3" {
		return true
	} else {
		return false
	}
}

func verify_n_wesolowski_with_b() bool {
	discriminant := "-0xd110d47ad62db12b68ef8328ba85f299408a7aeaa11875dce59434ed227c88d679ee9361a3f7518a9c4dc5cf2627ebd002192ecd9e93355e7e6b3cfc43758c4f"
	b_hex := "0x95bf9c93a498bbf1122214b04eeffba6eb6b7278b7bee30b4aece86a2d1813bbc1"
	x := "0200a0a6b0559653e0fefb1baea52e28fd8ffc14e834f1453742113006f667e44250cd928c5a063c1b2761c69b03c4648d2d0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	y_proof := "0100717922866a0cba4143d9082be10add54650b2c48745092135e1a501b597991252cf98f8d5d2c8c98b87707084b5304150100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	iters := 166670
	result := C.c_verify_n_wesolowski_with_b(C.CString(discriminant), C.CString(b_hex), C.CString(x), C.CString(y_proof), C.uint64_t(iters), C.uint64_t(0))

	defer C.free(unsafe.Pointer(result.y_from_compression))
	fmt.Println("verify_n_wesolowski_with_b", C._Bool(result.is_valid))
	fmt.Println("verify_n_wesolowski_with_b", C.GoString(result.y_from_compression))

	if C._Bool(result.is_valid) && C.GoString(result.y_from_compression) == "000012d3f71360994186a329d446bda961a56b7192b07be04834fe87c019a0e84224af8f7b45c7d3e5367ed52fe533fc6f4f0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" {
		return true
	} else {
		return false
	}
}

func verify_wesolowski() bool {
	discriminant := "-0x85c9a7f4adea5e2f56e7c78bb945e42b2ca2bb3c47f89af2dbd0c20df3f2fe2d1df7bfcfcfa9f7457f975b3f4483c7d58ccb8837acf91bb4dacb840064a903e7"
	initial_el := "08000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	result_y := "00006a5f7baee3c4e63bb81cfcbced11c0cfde70c09b69d4c5a8f03b1c559459c012f3700d0a35bc3f054acd020c03af9b3f0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	proof := "01006d4ecbbddba8275dc01b2bbd74cf7697faeb3d99c3e915b90b8c36c1c93c283fef42a105442853bccbedca2297fec90f0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	iters := 1000000
	is_valid := C.c_verify_wesolowski(C.CString(discriminant), C.CString(initial_el), C.CString(result_y), C.CString(proof), C.uint64_t(iters))

	fmt.Println("is_valid:", C._Bool(is_valid))

	if C._Bool(is_valid) {
		return true
	} else {
		return false
	}
}

func verify_n_wesolowski() bool {
	discriminant := "-0xb606f6683afa834d33e59cd2ae68c2b8919ed0da7ffc3d61057000934ffb20dee319c95fdaeb0955ca3bd49c0a72c29c0286e0bd2e3fb488af154444ea88e1c7"
	initial_el := "08000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	recursion := "010029900dacbdc04e9ef35e0aac7854855b3d39d206c2293dc34d10bfad1d7bf72a3bab7fab850a0772e5c982fee3f5981b0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000039654e8b52d28044b16eee5377d50c3f486e8502dcee822f1fc25667b930495ec9e6f957e651c837fd91535495b6571101000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000028b0a97ba7040fe1b29afd58a6ca4281d12e2138122d101cd88b943e821ccb3be3f7e7902002a1c375c4134525852215fd201f4d6f2d93bb09345b887bf73a18cbdf163c034cde44754cc5522158f956430567f070101000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000028b0a9b2d93fb4d57639983bd7486b7f610107482bdbc786f2154ec2e79fe9aa0bd5bbd0100cdb2dee0d01a8056f89c91ba88afad89707a61304bbcd07fc3523f6ee252fc0451192e43bc478207bd525287030c680301000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000028b0ab58bdd3ff8361736ac6dcfc1bb1b57f6f4c68a719828e607a0db0debc3e56af1c50300605b70a30c0ef198c47d9cc27e6f9dec6c9f39f3029fa118c26ed035e1feb9278b62941a6c952b6f58a5ecac59675c4f01000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000028b0ace46d722343167fcac019102e40a29daa98ded3ef70c7ab5ba82e733938062696f0300e7c86a740c0c4e863c87daf8deb575b9d8157b15ff70c90e31026eafe9296c1030ab2b2fd062f61f4a2ec498aae0042203020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000028b0acab16baa722712450bd389a0b8237f26dba2512aa49f8dbcc2f86439bc567849d30300eac6285ddb24cb165d61dbd1ccea99416d458299a1a3c050d704780e677ea00d15364c64c2fdf138950a79c415d6df0d0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	iters := 1000000
	discriminant_size := 512
	is_valid := C.c_verify_n_wesolowski(C.CString(discriminant), C.CString(initial_el), C.CString(recursion), C.uint64_t(iters), C.uint64_t(discriminant_size), C.uint64_t(5))

	fmt.Println("is_valid:", C._Bool(is_valid))

	if C._Bool(is_valid) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(create_discriminant())
	fmt.Println(prove())
	fmt.Println(get_b_from_n_wesolowski())
	fmt.Println(verify_n_wesolowski_with_b())

	fmt.Println(verify_wesolowski())
	fmt.Println(verify_n_wesolowski())
}