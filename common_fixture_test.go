package oak_test

import "github.com/phogolabs/parcello"

func addResourceWithMissingMigrationDir(manager *parcello.Manager) {
	manager.Add([]byte{
		31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 42, 78, 46, 202, 44, 40,
		209, 47, 169, 40, 209, 43, 169, 40, 97, 160, 9, 48, 48, 48,
		48, 48, 51, 49, 1, 211, 230, 102, 166, 96, 218, 192, 8, 194,
		135, 2, 6, 67, 99, 35, 115, 3, 67, 67, 19, 83, 99, 67, 6,
		3, 67, 19, 35, 3, 67, 6, 5, 3, 218, 56, 7, 21, 148, 22, 151,
		36, 22, 49, 24, 24, 20, 37, 230, 36, 103, 224, 81, 87, 92,
		146, 152, 150, 134, 71, 30, 230, 17, 24, 61, 68, 0, 0, 0,
		0, 255, 255,
	})
}

func addResourceWithMissingMigrations(manager *parcello.Manager) {
	manager.Add([]byte{
		31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 236, 81, 97, 139, 218,
		64, 16, 205, 231, 252, 138, 249, 168, 197, 173, 147, 53, 217,
		20, 161, 72, 90, 183, 84, 176, 90, 204, 218, 126, 148, 37,
		89, 109, 32, 137, 233, 238, 230, 255, 151, 68, 211, 11, 119,
		168, 31, 60, 238, 56, 184, 71, 96, 178, 179, 243, 222, 190,
		225, 21, 217, 65, 75, 155, 29, 203, 49, 34, 50, 12, 168, 143,
		45, 118, 70, 217, 186, 250, 104, 254, 230, 206, 221, 104,
		244, 152, 239, 183, 53, 100, 65, 91, 145, 158, 206, 103, 56,
		222, 132, 134, 136, 1, 69, 164, 14, 122, 140, 161, 231, 0,
		222, 255, 244, 109, 212, 198, 74, 237, 32, 106, 153, 39, 127,
		174, 204, 25, 43, 247, 251, 43, 247, 221, 34, 93, 125, 35,
		48, 137, 206, 42, 59, 78, 142, 69, 33, 203, 244, 121, 242,
		126, 140, 91, 249, 7, 33, 235, 242, 247, 194, 166, 239, 5,
		136, 236, 61, 255, 151, 0, 33, 16, 213, 246, 72, 14, 170,
		84, 90, 90, 149, 130, 180, 32, 106, 5, 81, 165, 129, 250,
		224, 177, 105, 48, 153, 210, 16, 190, 196, 2, 40, 122, 159,
		92, 151, 16, 40, 101, 161, 166, 96, 84, 174, 18, 75, 100,
		158, 147, 218, 40, 109, 8, 117, 99, 190, 228, 95, 5, 124,
		128, 111, 155, 245, 15, 104, 187, 79, 9, 77, 251, 194, 236,
		239, 239, 124, 195, 33, 75, 225, 51, 204, 122, 196, 172, 52,
		74, 255, 39, 46, 86, 49, 223, 8, 88, 172, 196, 250, 68, 131,
		65, 150, 142, 96, 159, 105, 99, 119, 13, 97, 4, 185, 60, 255,
		14, 221, 95, 209, 114, 203, 99, 24, 204, 70, 208, 124, 195,
		158, 108, 93, 165, 210, 170, 78, 118, 251, 115, 30, 9, 126,
		54, 18, 115, 209, 19, 108, 236, 244, 68, 91, 119, 23, 172,
		166, 42, 87, 15, 154, 115, 190, 228, 130, 95, 220, 240, 181,
		211, 119, 156, 127, 0, 0, 0, 255, 255,
	})
}

func addResource(manager *parcello.Manager) {
	manager.Add([]byte{
		31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 236, 149, 109, 79, 219,
		48, 16, 199, 251, 218, 159, 226, 94, 194, 212, 12, 59, 15,
		14, 20, 33, 20, 192, 104, 213, 250, 164, 196, 108, 240, 170,
		242, 26, 67, 35, 165, 105, 103, 59, 218, 215, 159, 236, 210,
		54, 176, 2, 147, 138, 96, 147, 248, 171, 146, 219, 248, 238,
		114, 87, 231, 247, 207, 172, 184, 83, 194, 20, 243, 234, 0,
		99, 76, 113, 228, 135, 216, 105, 172, 165, 169, 23, 159, 245,
		207, 178, 181, 179, 108, 61, 26, 134, 110, 141, 105, 228,
		86, 236, 47, 127, 91, 133, 52, 106, 145, 192, 167, 148, 196,
		36, 160, 184, 133, 9, 165, 1, 109, 1, 222, 253, 214, 47, 171,
		214, 70, 168, 22, 198, 74, 148, 147, 233, 51, 113, 218, 136,
		219, 219, 103, 246, 239, 103, 89, 175, 255, 137, 60, 15, 146,
		218, 204, 189, 59, 89, 73, 37, 140, 204, 65, 24, 224, 211,
		26, 146, 133, 2, 114, 4, 228, 168, 19, 248, 29, 114, 8, 231,
		44, 227, 224, 99, 114, 136, 60, 15, 70, 165, 20, 90, 66, 62,
		135, 106, 110, 96, 50, 21, 213, 157, 4, 51, 149, 80, 137,
		153, 4, 97, 140, 42, 126, 212, 70, 106, 100, 131, 237, 181,
		14, 212, 11, 132, 206, 83, 150, 112, 6, 60, 57, 235, 49, 232,
		94, 194, 96, 200, 129, 93, 119, 51, 158, 193, 250, 57, 212,
		176, 135, 160, 200, 97, 45, 206, 174, 249, 242, 155, 13, 31,
		92, 245, 122, 48, 74, 187, 253, 36, 189, 129, 175, 236, 166,
		141, 32, 151, 122, 162, 138, 133, 77, 222, 18, 220, 70, 48,
		81, 210, 78, 54, 22, 6, 128, 119, 251, 44, 227, 73, 127, 180,
		14, 64, 251, 199, 155, 46, 243, 249, 175, 10, 161, 139, 116,
		56, 218, 116, 249, 71, 135, 199, 239, 125, 104, 175, 168,
		13, 255, 246, 108, 113, 136, 41, 57, 194, 152, 68, 227, 90,
		75, 165, 223, 134, 255, 32, 14, 29, 255, 62, 141, 3, 28, 7,
		142, 127, 250, 193, 255, 155, 104, 43, 255, 151, 170, 112,
		252, 3, 181, 252, 99, 220, 33, 209, 107, 240, 255, 0, 127,
		247, 124, 89, 214, 45, 236, 221, 1, 111, 66, 221, 132, 23,
		110, 11, 165, 205, 216, 21, 118, 116, 55, 247, 74, 209, 220,
		178, 36, 63, 66, 121, 43, 201, 238, 214, 199, 8, 189, 247,
		127, 255, 47, 104, 105, 157, 7, 147, 249, 108, 38, 170, 252,
		117, 120, 127, 172, 151, 248, 143, 3, 223, 242, 31, 99, 28,
		17, 236, 91, 254, 195, 56, 198, 31, 252, 191, 133, 182, 191,
		255, 107, 233, 248, 247, 67, 32, 180, 19, 5, 29, 63, 134,
		179, 21, 254, 27, 194, 180, 44, 229, 196, 120, 162, 44, 61,
		135, 20, 202, 88, 143, 157, 115, 248, 4, 151, 233, 176, 191,
		196, 108, 231, 112, 123, 121, 107, 228, 247, 47, 44, 101,
		214, 60, 78, 224, 116, 247, 180, 162, 210, 82, 221, 167, 117,
		7, 25, 75, 185, 53, 165, 225, 202, 166, 138, 188, 221, 48,
		162, 246, 198, 120, 246, 209, 183, 164, 119, 197, 50, 216,
		59, 109, 131, 253, 236, 63, 240, 188, 92, 24, 185, 44, 122,
		53, 186, 176, 230, 183, 154, 156, 55, 125, 237, 196, 102,
		110, 188, 204, 118, 246, 68, 155, 185, 44, 229, 170, 226,
		5, 235, 49, 206, 158, 156, 237, 239, 206, 255, 55, 0, 0, 0,
		255, 255,
	})
}