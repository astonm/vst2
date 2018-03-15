package vst2

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	samples32 = [][]float32{
		{
			-0.0027160645, -0.0039978027, -0.0071411133, -0.0065307617, 0.0038757324, 0.021972656, 0.041229248, 0.055511475, 0.064971924, 0.07342529, 0.08300781, 0.092681885, 0.10070801, 0.110809326, 0.12677002, 0.15231323, 0.19058228, 0.24459839, 0.3140869, 0.38861084, 0.44683838, 0.47177124, 0.46643066, 0.45007324, 0.4449768, 0.45724487, 0.47451782, 0.48321533, 0.47824097, 0.46679688, 0.45999146, 0.46765137, 0.491333, 0.52505493, 0.555542, 0.57055664, 0.5701599, 0.56591797, 0.5706787, 0.58740234, 0.60510254, 0.6090698, 0.5979004, 0.5837097, 0.58288574, 0.6016846, 0.63098145, 0.6526184, 0.6595459, 0.6639404, 0.6861267, 0.73825073, 0.8117676, 0.8830261, 0.9341736, 0.9536133, 0.95010376, 0.9467163, 0.9489136, 0.94799805, 0.9470825, 0.9473877, 0.9465332, 0.9461365, 0.9458313, 0.94555664, 0.94525146, 0.94454956, 0.94351196, 0.9433899, 0.9428711, 0.94226074, 0.94226074, 0.9415283, 0.94104004, 0.9407959, 0.93963623, 0.9390869, 0.93881226, 0.9384155, 0.93777466, 0.9367676, 0.93652344, 0.9362488, 0.9352417, 0.9342041, 0.93414307, 0.93356323, 0.93304443, 0.9326782, 0.93182373, 0.93136597, 0.93084717, 0.9299927, 0.92926025, 0.928833, 0.9283447, 0.9275818, 0.9270935, 0.9267273, 0.92578125, 0.92526245, 0.92489624, 0.9238281, 0.9232178, 0.92266846, 0.9219055, 0.92123413, 0.92123413, 0.9200134, 0.9194336, 0.9192505, 0.91848755, 0.9177246, 0.91744995, 0.9169922, 0.9158325, 0.9154968, 0.9150696, 0.91430664, 0.9137268, 0.91326904, 0.91275024, 0.91189575, 0.91137695, 0.9107971, 0.9099121, 0.9095154, 0.9087219, 0.9081421, 0.9077759, 0.9072571, 0.9066162, 0.90618896, 0.9053345, 0.9050293, 0.9039612, 0.9034729, 0.9031677, 0.90237427, 0.9014282, 0.90130615, 0.90078735, 0.89974976, 0.89959717, 0.8988342, 0.8982239, 0.8980713, 0.8967285, 0.89675903, 0.895813, 0.8949585, 0.89016724, 0.8718872, 0.8529358, 0.85061646, 0.8433838, 0.81521606, 0.78323364, 0.7631836, 0.7581787, 0.7541809, 0.7345276, 0.7017517, 0.6737976, 0.66192627, 0.6621704, 0.65826416, 0.6414795, 0.6186218, 0.59976196, 0.58813477, 0.5793457, 0.5661621, 0.54629517, 0.5232239, 0.49957275, 0.47854614, 0.4595642, 0.44082642, 0.41912842, 0.39367676, 0.3659668, 0.34313965, 0.32885742, 0.32141113, 0.31307983, 0.29656982, 0.2741394, 0.2564392, 0.24880981, 0.2480774, 0.24499512, 0.23535156, 0.2230835, 0.21395874, 0.21069336, 0.20672607, 0.19671631, 0.18139648, 0.16766357, 0.16088867, 0.16131592, 0.16290283, 0.15875244, 0.14880371, 0.14047241, 0.1375122, 0.13754272, 0.1312561, 0.11392212, 0.087524414, 0.06283569, 0.044067383, 0.028747559, 0.013336182, -0.002105713, -0.014251709, -0.022949219, -0.033081055, -0.050811768 - 0.07687378, -0.107299805, -0.13864136, -0.16897583, -0.19692993, -0.21844482, -0.23138428, -0.23873901, -0.24871826, -0.26620483, -0.28753662, -0.30618286, -0.3239441, -0.3468628, -0.37780762, -0.40966797, -0.4310913, -0.4385376, -0.43984985, -0.44317627, -0.4505005, -0.45721436, -0.46203613, -0.47052002, -0.49023438, -0.51968384, -0.5514221, -0.57769775, -0.5969238, -0.61032104, -0.6180115, -0.6163635, -0.6074829, -0.59988403, -0.60150146, -0.6123352, -0.62579346, -0.6381836, -0.65078735, -0.66641235, -0.682251, -0.69226074, -0.6951904, -0.697876, -0.70532227, -0.7165222, -0.724823, -0.7284851, -0.73114014, -0.73703, -0.744751, -0.7517395, -0.7571411, -0.7656555, -0.77575684, -0.78308105, -0.783905, -0.7837219, -0.7885742, -0.79852295, -0.8067322, -0.80700684, -0.8046875, -0.80966187, -0.8235779, -0.8360596, -0.836853, -0.82455444, -0.8106384, -0.80633545, -0.8121643, -0.82003784, -0.82110596, -0.8159485, -0.8136902, -0.8216858, -0.8374634, -0.8496399, -0.8527527, -0.8534546, -0.8637085, -0.88290405, -0.8952637, -0.88739014, -0.87106323, -0.8749695, -0.908844, -0.9442749, -0.94314575, -0.9057007, -0.8771057, -0.89553833, -0.9460144, -0.9734802, -0.95080566, -0.9076538, -0.8945923, -0.9199524, -0.94784546, -0.95092773, -0.9447632, -0.9534912, -0.9665222, -0.9493103, -0.90164185, -0.87387085, -0.90740967, -0.9699402, -0.98379517, -0.92019653, -0.84490967, -0.8469238, -0.9286194, -0.99887085, -0.9760437, -0.88500977, -0.82318115, -0.8460083, -0.90844727, -0.9246521, -0.8715515, -0.8070984, -0.7927246, -0.8227844, -0.8427124, -0.8262329, -0.801239, -0.8036194, -0.8232422, -0.8187561, -0.7727051, -0.7202759, -0.707428, -0.73535156, -0.7595215, -0.74105835, -0.69088745, -0.651947, -0.6512146, -0.67611694, -0.6911621, -0.67507935, -0.6367798, -0.60113525, -0.58914185, -0.602478, -0.6260681, -0.6374817, -0.62338257, -0.59176636, -0.5643921, -0.55389404, -0.5558777, -0.5494995, -0.52194214, -0.4753418, -0.42858887, -0.40090942, -0.40319824, -0.4307251 - 0.46392822, -0.4769287, -0.4588318, -0.42556763, -0.40551758, -0.40753174, -0.40896606, -0.3852539, -0.3413086, -0.31314087, -0.32635498, -0.36328125, -0.3829956, -0.3668213, -0.337677, -0.32650757, -0.3361206, -0.34222412, -0.33035278, -0.3109131, -0.29956055, -0.2885437, -0.2569275, -0.20565796, -0.1643982, -0.15802002, -0.17419434, -0.17623901, -0.14611816, -0.10662842, -0.09020996, -0.09765625, -0.09631348, -0.065216064, -0.0184021, 0.008453369, 0.0012207031, -0.017486572, -0.015106201, 0.014923096, 0.051605225, 0.075408936, 0.087524414, 0.10205078, 0.12762451, 0.1609497, 0.19122314, 0.2098999, 0.21704102, 0.21661377, 0.21902466, 0.23583984, 0.26986694, 0.31033325, 0.34545898, 0.37139893, 0.39263916, 0.41351318, 0.43130493, 0.44195557, 0.4508667, 0.46948242, 0.4954834, 0.5108032, 0.5023804, 0.48156738, 0.47988892, 0.51226807, 0.5572815, 0.58102417, 0.5742798, 0.559845, 0.56741333, 0.60110474, 0.64263916, 0.67437744, 0.6965332, 0.7150574, 0.7319641, 0.74438477, 0.75839233, 0.78186035, 0.81207275, 0.8309021, 0.8224182, 0.79211426, 0.7640991, 0.75460815, 0.7586365, 0.7642822, 0.7694092, 0.78253174, 0.8069763, 0.8355713, 0.86026, 0.8809509, 0.9004822, 0.9154663, 0.91726685, 0.90792847, 0.9007263, 0.90896606, 0.9289856, 0.9402771, 0.93685913, 0.93447876, 0.93066406, 0.92163086, 0.9219971, 0.93292236, 0.93548584, 0.9326172, 0.932312, 0.9321594, 0.93078613, 0.9276428, 0.9147949, 0.8964844, 0.8874512, 0.8934021, 0.91033936, 0.92663574, 0.92926025, 0.9249573, 0.92544556, 0.9252014, 0.923645, 0.92471313, 0.9231262, 0.9221802, 0.9222717, 0.9221802, 0.9222717,
		}, {
			-0.0032653809, -0.0038146973, -0.0069274902, -0.0077209473, 0.00048828125, 0.017669678, 0.03768921, 0.053497314, 0.06427002, 0.07284546, 0.08291626, 0.092926025, 0.101501465, 0.11090088, 0.12573242, 0.1496582, 0.18566895, 0.23712158, 0.3048401, 0.38153076, 0.4463501, 0.48077393, 0.48117065, 0.4656067, 0.4569397, 0.46606445, 0.48397827, 0.4953003, 0.49295044, 0.4814453, 0.47244263, 0.47601318, 0.49642944, 0.5291443, 0.561615, 0.58081055, 0.58325195, 0.5786743, 0.5803528, 0.5953064, 0.61380005, 0.62197876, 0.6130676, 0.5973816, 0.5923462, 0.6072998, 0.63586426, 0.6604004, 0.6701355, 0.6734314, 0.6902466, 0.7367554, 0.8083801, 0.8835144, 0.94137573, 0.9691467, 0.9689636, 0.96417236, 0.96655273, 0.9662781, 0.96432495, 0.9647522, 0.9637146, 0.9630127, 0.9622803, 0.96157837, 0.96084595, 0.9597473, 0.95843506, 0.9577637, 0.9569092, 0.95584106, 0.9551697, 0.95401, 0.95358276, 0.9526367, 0.9515991, 0.9507446, 0.950531, 0.9495239, 0.9484863, 0.9475403, 0.9468994, 0.94628906, 0.94537354, 0.9440918, 0.9439697, 0.9430237, 0.94262695, 0.9417114, 0.9409485, 0.94039917, 0.9393921, 0.93878174, 0.9378052, 0.93743896, 0.93670654, 0.93603516, 0.9352112, 0.9350281, 0.9338684, 0.9333191, 0.9328003, 0.9320679, 0.9311218, 0.9306946, 0.9298706, 0.9291382, 0.928772, 0.9281616, 0.92684937, 0.9267578, 0.92645264, 0.9254761, 0.9249878, 0.9244385, 0.923584, 0.92318726, 0.92266846, 0.92178345, 0.92123413, 0.9208679, 0.920105, 0.91918945, 0.9187622, 0.91799927, 0.9173279, 0.91656494, 0.9161377, 0.91571045, 0.9152527, 0.9144592, 0.9137573, 0.91360474, 0.9128418, 0.9119568, 0.9116211, 0.9109497, 0.9105835, 0.90979004, 0.9088135, 0.9085388, 0.90826416, 0.90704346, 0.90670776, 0.9062195, 0.90548706, 0.9052124, 0.9043274, 0.9038391, 0.9033203, 0.90231323, 0.8996887, 0.88360596, 0.8621826, 0.85672, 0.85317993, 0.8277588, 0.79437256, 0.7703552, 0.76242065, 0.7597046, 0.74334717, 0.71154785, 0.68078613, 0.664978, 0.66293335, 0.6612854, 0.6468506, 0.6239319, 0.60305786, 0.5899658, 0.58102417, 0.5687561, 0.5496826, 0.5267334, 0.50283813, 0.48046875, 0.46096802, 0.44177246, 0.42077637, 0.39544678, 0.36727905, 0.3423462, 0.3260498, 0.31741333, 0.30975342, 0.29489136, 0.27285767, 0.2532959, 0.24353027, 0.24206543, 0.24023438, 0.23181152, 0.21957397, 0.2093811, 0.20526123, 0.20227051, 0.19396973, 0.17938232, 0.16485596, 0.15646362, 0.15570068, 0.15789795, 0.15542603, 0.14648438, 0.13723755, 0.13348389, 0.13369751, 0.1295166, 0.11419678, 0.08932495, 0.063079834, 0.042816162, 0.026885986, 0.011383057, -0.004638672, -0.01776123, -0.02746582, -0.036499023, -0.052703857, -0.07720947, -0.10772705, -0.13937378, -0.17102051, -0.2001648, -0.22399902, -0.23895264, -0.2470398, -0.25634766, -0.27227783, -0.29360962, -0.3133545, -0.33096313, -0.35238647, -0.3824768, -0.41558838, -0.44033813, -0.4508667, -0.45230103, -0.45483398, -0.46157837, -0.4685669, -0.4732666, -0.48025513, -0.49716187, -0.5252991, -0.5574341, -0.58566284, -0.6065674, -0.62142944, -0.6305847, -0.6309204, -0.62286377, -0.6138611, -0.6127014, -0.6218262, -0.63516235, -0.6473999, -0.65982056, -0.67489624, -0.69122314, -0.7029114, -0.70687866, -0.7088928, -0.71536255, -0.72610474, -0.7352295, -0.73950195, -0.7420044, -0.74679565, -0.7546997, -0.76174927, -0.767395, -0.77474976, -0.78500366, -0.7932434, -0.7953491, -0.79422, -0.7979126, -0.8074341, -0.8163147, -0.8180847, -0.81555176, -0.81811523, -0.83065796, -0.8450012, -0.84851074, -0.83810425, -0.8231201, -0.8161011, -0.8200073, -0.8280945, -0.83084106, -0.82650757, -0.82244873, -0.8278198, -0.8430176, -0.8569641, -0.8617554, -0.861969, -0.8698425, -0.888031, -0.90393066, -0.90045166, -0.88357544, -0.8807373, -0.9099121, -0.94924927, -0.9577942, -0.92510986, -0.88974, -0.8977661, -0.9454346, -0.9822998, -0.96972656, -0.9260864, -0.9031677, -0.92251587, -0.9532776, -0.96203613, -0.9551697, -0.96029663, -0.97540283, -0.96640015, -0.9215393, -0.8838196, -0.90393066, -0.96713257, -0.9972534, -0.9467468, -0.8647766, -0.8450012, -0.9161377, -0.9987793, -0.9981384, -0.91345215, -0.83639526, -0.84161377, -0.90338135, -0.934906, -0.89312744, -0.823822, -0.79611206, -0.8203125, -0.84677124, -0.8371887, -0.8097229, -0.8052063, -0.8243408, -0.8285217, -0.78933716, -0.73309326, -0.70877075, -0.7309265, -0.76031494, -0.75177, -0.7045288, -0.658905, -0.6498108, -0.67193604, -0.6920471, -0.68273926, -0.6465149, -0.60821533, -0.5895691, -0.59832764, -0.6218567, -0.6378174 - 0.629303, -0.59973145, -0.56918335, -0.5548401, -0.55548096, -0.55267334, -0.5291443, -0.48501587, -0.43563843, -0.4019165, -0.39712524, -0.42053223, -0.45510864, -0.47479248, -0.4630127, -0.43029785, -0.40567017, -0.40408325, -0.40847778, -0.39056396, -0.34796143, -0.31292725, -0.31723022, -0.35290527, -0.37963867, -0.37069702, -0.341156, -0.32455444, -0.33151245, -0.34109497, -0.33261108 - 0.31311035, -0.2998352, -0.29040527, -0.26385498, -0.21435547, -0.16738892, -0.15228271, -0.1666565, -0.17404175, -0.14981079, -0.108795166, -0.085754395, -0.08950806, -0.09338379, -0.06842041, -0.02130127, 0.012908936, 0.012054443, -0.0072631836, -0.010955811, 0.014343262, 0.052001953, 0.07937622, 0.093322754, 0.10632324, 0.12997437, 0.16271973, 0.1949768, 0.21697998, 0.22683716, 0.22729492, 0.22787476, 0.24108887, 0.27175903, 0.31277466, 0.35046387, 0.37854004, 0.40048218, 0.42190552, 0.44143677, 0.45367432, 0.4619751, 0.47805786, 0.50439453, 0.5238037, 0.5198059, 0.4991455, 0.49105835, 0.51657104, 0.5618286, 0.59280396, 0.5909729, 0.5753784, 0.57644653, 0.60598755, 0.6477661, 0.6826477, 0.7065735, 0.72592163, 0.7437744, 0.7574768, 0.7702942, 0.79159546, 0.82211304, 0.84536743, 0.84317017, 0.81503296, 0.783844, 0.7694092, 0.7713318, 0.7767029, 0.7810974, 0.7915039, 0.8138428, 0.84265137, 0.8684082, 0.89016724, 0.9100342, 0.9269409, 0.9324341, 0.92440796, 0.91519165, 0.91955566, 0.93881226, 0.9534607, 0.9517212, 0.9483032, 0.94589233, 0.9362793, 0.9331665, 0.94351196, 0.9481201, 0.9451904, 0.9439087, 0.94403076, 0.9426575, 0.9401245, 0.9293823, 0.91033936, 0.8980408, 0.90045166, 0.9158325, 0.9336548, 0.9397583, 0.93515015, 0.93460083, 0.93536377, 0.93325806, 0.9333801, 0.9326172, 0.93118286, 0.9313965, 0.93118286, 0.9313965,
		},
	}
	samples64 = [][]float64{
		{
			-0.0027160645, -0.0039978027, -0.0071411133, -0.0065307617, 0.0038757324, 0.021972656, 0.041229248, 0.055511475, 0.064971924, 0.07342529, 0.08300781, 0.092681885, 0.10070801, 0.110809326, 0.12677002, 0.15231323, 0.19058228, 0.24459839, 0.3140869, 0.38861084, 0.44683838, 0.47177124, 0.46643066, 0.45007324, 0.4449768, 0.45724487, 0.47451782, 0.48321533, 0.47824097, 0.46679688, 0.45999146, 0.46765137, 0.491333, 0.52505493, 0.555542, 0.57055664, 0.5701599, 0.56591797, 0.5706787, 0.58740234, 0.60510254, 0.6090698, 0.5979004, 0.5837097, 0.58288574, 0.6016846, 0.63098145, 0.6526184, 0.6595459, 0.6639404, 0.6861267, 0.73825073, 0.8117676, 0.8830261, 0.9341736, 0.9536133, 0.95010376, 0.9467163, 0.9489136, 0.94799805, 0.9470825, 0.9473877, 0.9465332, 0.9461365, 0.9458313, 0.94555664, 0.94525146, 0.94454956, 0.94351196, 0.9433899, 0.9428711, 0.94226074, 0.94226074, 0.9415283, 0.94104004, 0.9407959, 0.93963623, 0.9390869, 0.93881226, 0.9384155, 0.93777466, 0.9367676, 0.93652344, 0.9362488, 0.9352417, 0.9342041, 0.93414307, 0.93356323, 0.93304443, 0.9326782, 0.93182373, 0.93136597, 0.93084717, 0.9299927, 0.92926025, 0.928833, 0.9283447, 0.9275818, 0.9270935, 0.9267273, 0.92578125, 0.92526245, 0.92489624, 0.9238281, 0.9232178, 0.92266846, 0.9219055, 0.92123413, 0.92123413, 0.9200134, 0.9194336, 0.9192505, 0.91848755, 0.9177246, 0.91744995, 0.9169922, 0.9158325, 0.9154968, 0.9150696, 0.91430664, 0.9137268, 0.91326904, 0.91275024, 0.91189575, 0.91137695, 0.9107971, 0.9099121, 0.9095154, 0.9087219, 0.9081421, 0.9077759, 0.9072571, 0.9066162, 0.90618896, 0.9053345, 0.9050293, 0.9039612, 0.9034729, 0.9031677, 0.90237427, 0.9014282, 0.90130615, 0.90078735, 0.89974976, 0.89959717, 0.8988342, 0.8982239, 0.8980713, 0.8967285, 0.89675903, 0.895813, 0.8949585, 0.89016724, 0.8718872, 0.8529358, 0.85061646, 0.8433838, 0.81521606, 0.78323364, 0.7631836, 0.7581787, 0.7541809, 0.7345276, 0.7017517, 0.6737976, 0.66192627, 0.6621704, 0.65826416, 0.6414795, 0.6186218, 0.59976196, 0.58813477, 0.5793457, 0.5661621, 0.54629517, 0.5232239, 0.49957275, 0.47854614, 0.4595642, 0.44082642, 0.41912842, 0.39367676, 0.3659668, 0.34313965, 0.32885742, 0.32141113, 0.31307983, 0.29656982, 0.2741394, 0.2564392, 0.24880981, 0.2480774, 0.24499512, 0.23535156, 0.2230835, 0.21395874, 0.21069336, 0.20672607, 0.19671631, 0.18139648, 0.16766357, 0.16088867, 0.16131592, 0.16290283, 0.15875244, 0.14880371, 0.14047241, 0.1375122, 0.13754272, 0.1312561, 0.11392212, 0.087524414, 0.06283569, 0.044067383, 0.028747559, 0.013336182, -0.002105713, -0.014251709, -0.022949219, -0.033081055, -0.050811768 - 0.07687378, -0.107299805, -0.13864136, -0.16897583, -0.19692993, -0.21844482, -0.23138428, -0.23873901, -0.24871826, -0.26620483, -0.28753662, -0.30618286, -0.3239441, -0.3468628, -0.37780762, -0.40966797, -0.4310913, -0.4385376, -0.43984985, -0.44317627, -0.4505005, -0.45721436, -0.46203613, -0.47052002, -0.49023438, -0.51968384, -0.5514221, -0.57769775, -0.5969238, -0.61032104, -0.6180115, -0.6163635, -0.6074829, -0.59988403, -0.60150146, -0.6123352, -0.62579346, -0.6381836, -0.65078735, -0.66641235, -0.682251, -0.69226074, -0.6951904, -0.697876, -0.70532227, -0.7165222, -0.724823, -0.7284851, -0.73114014, -0.73703, -0.744751, -0.7517395, -0.7571411, -0.7656555, -0.77575684, -0.78308105, -0.783905, -0.7837219, -0.7885742, -0.79852295, -0.8067322, -0.80700684, -0.8046875, -0.80966187, -0.8235779, -0.8360596, -0.836853, -0.82455444, -0.8106384, -0.80633545, -0.8121643, -0.82003784, -0.82110596, -0.8159485, -0.8136902, -0.8216858, -0.8374634, -0.8496399, -0.8527527, -0.8534546, -0.8637085, -0.88290405, -0.8952637, -0.88739014, -0.87106323, -0.8749695, -0.908844, -0.9442749, -0.94314575, -0.9057007, -0.8771057, -0.89553833, -0.9460144, -0.9734802, -0.95080566, -0.9076538, -0.8945923, -0.9199524, -0.94784546, -0.95092773, -0.9447632, -0.9534912, -0.9665222, -0.9493103, -0.90164185, -0.87387085, -0.90740967, -0.9699402, -0.98379517, -0.92019653, -0.84490967, -0.8469238, -0.9286194, -0.99887085, -0.9760437, -0.88500977, -0.82318115, -0.8460083, -0.90844727, -0.9246521, -0.8715515, -0.8070984, -0.7927246, -0.8227844, -0.8427124, -0.8262329, -0.801239, -0.8036194, -0.8232422, -0.8187561, -0.7727051, -0.7202759, -0.707428, -0.73535156, -0.7595215, -0.74105835, -0.69088745, -0.651947, -0.6512146, -0.67611694, -0.6911621, -0.67507935, -0.6367798, -0.60113525, -0.58914185, -0.602478, -0.6260681, -0.6374817, -0.62338257, -0.59176636, -0.5643921, -0.55389404, -0.5558777, -0.5494995, -0.52194214, -0.4753418, -0.42858887, -0.40090942, -0.40319824, -0.4307251 - 0.46392822, -0.4769287, -0.4588318, -0.42556763, -0.40551758, -0.40753174, -0.40896606, -0.3852539, -0.3413086, -0.31314087, -0.32635498, -0.36328125, -0.3829956, -0.3668213, -0.337677, -0.32650757, -0.3361206, -0.34222412, -0.33035278, -0.3109131, -0.29956055, -0.2885437, -0.2569275, -0.20565796, -0.1643982, -0.15802002, -0.17419434, -0.17623901, -0.14611816, -0.10662842, -0.09020996, -0.09765625, -0.09631348, -0.065216064, -0.0184021, 0.008453369, 0.0012207031, -0.017486572, -0.015106201, 0.014923096, 0.051605225, 0.075408936, 0.087524414, 0.10205078, 0.12762451, 0.1609497, 0.19122314, 0.2098999, 0.21704102, 0.21661377, 0.21902466, 0.23583984, 0.26986694, 0.31033325, 0.34545898, 0.37139893, 0.39263916, 0.41351318, 0.43130493, 0.44195557, 0.4508667, 0.46948242, 0.4954834, 0.5108032, 0.5023804, 0.48156738, 0.47988892, 0.51226807, 0.5572815, 0.58102417, 0.5742798, 0.559845, 0.56741333, 0.60110474, 0.64263916, 0.67437744, 0.6965332, 0.7150574, 0.7319641, 0.74438477, 0.75839233, 0.78186035, 0.81207275, 0.8309021, 0.8224182, 0.79211426, 0.7640991, 0.75460815, 0.7586365, 0.7642822, 0.7694092, 0.78253174, 0.8069763, 0.8355713, 0.86026, 0.8809509, 0.9004822, 0.9154663, 0.91726685, 0.90792847, 0.9007263, 0.90896606, 0.9289856, 0.9402771, 0.93685913, 0.93447876, 0.93066406, 0.92163086, 0.9219971, 0.93292236, 0.93548584, 0.9326172, 0.932312, 0.9321594, 0.93078613, 0.9276428, 0.9147949, 0.8964844, 0.8874512, 0.8934021, 0.91033936, 0.92663574, 0.92926025, 0.9249573, 0.92544556, 0.9252014, 0.923645, 0.92471313, 0.9231262, 0.9221802, 0.9222717, 0.9221802, 0.9222717,
		}, {
			-0.0032653809, -0.0038146973, -0.0069274902, -0.0077209473, 0.00048828125, 0.017669678, 0.03768921, 0.053497314, 0.06427002, 0.07284546, 0.08291626, 0.092926025, 0.101501465, 0.11090088, 0.12573242, 0.1496582, 0.18566895, 0.23712158, 0.3048401, 0.38153076, 0.4463501, 0.48077393, 0.48117065, 0.4656067, 0.4569397, 0.46606445, 0.48397827, 0.4953003, 0.49295044, 0.4814453, 0.47244263, 0.47601318, 0.49642944, 0.5291443, 0.561615, 0.58081055, 0.58325195, 0.5786743, 0.5803528, 0.5953064, 0.61380005, 0.62197876, 0.6130676, 0.5973816, 0.5923462, 0.6072998, 0.63586426, 0.6604004, 0.6701355, 0.6734314, 0.6902466, 0.7367554, 0.8083801, 0.8835144, 0.94137573, 0.9691467, 0.9689636, 0.96417236, 0.96655273, 0.9662781, 0.96432495, 0.9647522, 0.9637146, 0.9630127, 0.9622803, 0.96157837, 0.96084595, 0.9597473, 0.95843506, 0.9577637, 0.9569092, 0.95584106, 0.9551697, 0.95401, 0.95358276, 0.9526367, 0.9515991, 0.9507446, 0.950531, 0.9495239, 0.9484863, 0.9475403, 0.9468994, 0.94628906, 0.94537354, 0.9440918, 0.9439697, 0.9430237, 0.94262695, 0.9417114, 0.9409485, 0.94039917, 0.9393921, 0.93878174, 0.9378052, 0.93743896, 0.93670654, 0.93603516, 0.9352112, 0.9350281, 0.9338684, 0.9333191, 0.9328003, 0.9320679, 0.9311218, 0.9306946, 0.9298706, 0.9291382, 0.928772, 0.9281616, 0.92684937, 0.9267578, 0.92645264, 0.9254761, 0.9249878, 0.9244385, 0.923584, 0.92318726, 0.92266846, 0.92178345, 0.92123413, 0.9208679, 0.920105, 0.91918945, 0.9187622, 0.91799927, 0.9173279, 0.91656494, 0.9161377, 0.91571045, 0.9152527, 0.9144592, 0.9137573, 0.91360474, 0.9128418, 0.9119568, 0.9116211, 0.9109497, 0.9105835, 0.90979004, 0.9088135, 0.9085388, 0.90826416, 0.90704346, 0.90670776, 0.9062195, 0.90548706, 0.9052124, 0.9043274, 0.9038391, 0.9033203, 0.90231323, 0.8996887, 0.88360596, 0.8621826, 0.85672, 0.85317993, 0.8277588, 0.79437256, 0.7703552, 0.76242065, 0.7597046, 0.74334717, 0.71154785, 0.68078613, 0.664978, 0.66293335, 0.6612854, 0.6468506, 0.6239319, 0.60305786, 0.5899658, 0.58102417, 0.5687561, 0.5496826, 0.5267334, 0.50283813, 0.48046875, 0.46096802, 0.44177246, 0.42077637, 0.39544678, 0.36727905, 0.3423462, 0.3260498, 0.31741333, 0.30975342, 0.29489136, 0.27285767, 0.2532959, 0.24353027, 0.24206543, 0.24023438, 0.23181152, 0.21957397, 0.2093811, 0.20526123, 0.20227051, 0.19396973, 0.17938232, 0.16485596, 0.15646362, 0.15570068, 0.15789795, 0.15542603, 0.14648438, 0.13723755, 0.13348389, 0.13369751, 0.1295166, 0.11419678, 0.08932495, 0.063079834, 0.042816162, 0.026885986, 0.011383057, -0.004638672, -0.01776123, -0.02746582, -0.036499023, -0.052703857, -0.07720947, -0.10772705, -0.13937378, -0.17102051, -0.2001648, -0.22399902, -0.23895264, -0.2470398, -0.25634766, -0.27227783, -0.29360962, -0.3133545, -0.33096313, -0.35238647, -0.3824768, -0.41558838, -0.44033813, -0.4508667, -0.45230103, -0.45483398, -0.46157837, -0.4685669, -0.4732666, -0.48025513, -0.49716187, -0.5252991, -0.5574341, -0.58566284, -0.6065674, -0.62142944, -0.6305847, -0.6309204, -0.62286377, -0.6138611, -0.6127014, -0.6218262, -0.63516235, -0.6473999, -0.65982056, -0.67489624, -0.69122314, -0.7029114, -0.70687866, -0.7088928, -0.71536255, -0.72610474, -0.7352295, -0.73950195, -0.7420044, -0.74679565, -0.7546997, -0.76174927, -0.767395, -0.77474976, -0.78500366, -0.7932434, -0.7953491, -0.79422, -0.7979126, -0.8074341, -0.8163147, -0.8180847, -0.81555176, -0.81811523, -0.83065796, -0.8450012, -0.84851074, -0.83810425, -0.8231201, -0.8161011, -0.8200073, -0.8280945, -0.83084106, -0.82650757, -0.82244873, -0.8278198, -0.8430176, -0.8569641, -0.8617554, -0.861969, -0.8698425, -0.888031, -0.90393066, -0.90045166, -0.88357544, -0.8807373, -0.9099121, -0.94924927, -0.9577942, -0.92510986, -0.88974, -0.8977661, -0.9454346, -0.9822998, -0.96972656, -0.9260864, -0.9031677, -0.92251587, -0.9532776, -0.96203613, -0.9551697, -0.96029663, -0.97540283, -0.96640015, -0.9215393, -0.8838196, -0.90393066, -0.96713257, -0.9972534, -0.9467468, -0.8647766, -0.8450012, -0.9161377, -0.9987793, -0.9981384, -0.91345215, -0.83639526, -0.84161377, -0.90338135, -0.934906, -0.89312744, -0.823822, -0.79611206, -0.8203125, -0.84677124, -0.8371887, -0.8097229, -0.8052063, -0.8243408, -0.8285217, -0.78933716, -0.73309326, -0.70877075, -0.7309265, -0.76031494, -0.75177, -0.7045288, -0.658905, -0.6498108, -0.67193604, -0.6920471, -0.68273926, -0.6465149, -0.60821533, -0.5895691, -0.59832764, -0.6218567, -0.6378174 - 0.629303, -0.59973145, -0.56918335, -0.5548401, -0.55548096, -0.55267334, -0.5291443, -0.48501587, -0.43563843, -0.4019165, -0.39712524, -0.42053223, -0.45510864, -0.47479248, -0.4630127, -0.43029785, -0.40567017, -0.40408325, -0.40847778, -0.39056396, -0.34796143, -0.31292725, -0.31723022, -0.35290527, -0.37963867, -0.37069702, -0.341156, -0.32455444, -0.33151245, -0.34109497, -0.33261108 - 0.31311035, -0.2998352, -0.29040527, -0.26385498, -0.21435547, -0.16738892, -0.15228271, -0.1666565, -0.17404175, -0.14981079, -0.108795166, -0.085754395, -0.08950806, -0.09338379, -0.06842041, -0.02130127, 0.012908936, 0.012054443, -0.0072631836, -0.010955811, 0.014343262, 0.052001953, 0.07937622, 0.093322754, 0.10632324, 0.12997437, 0.16271973, 0.1949768, 0.21697998, 0.22683716, 0.22729492, 0.22787476, 0.24108887, 0.27175903, 0.31277466, 0.35046387, 0.37854004, 0.40048218, 0.42190552, 0.44143677, 0.45367432, 0.4619751, 0.47805786, 0.50439453, 0.5238037, 0.5198059, 0.4991455, 0.49105835, 0.51657104, 0.5618286, 0.59280396, 0.5909729, 0.5753784, 0.57644653, 0.60598755, 0.6477661, 0.6826477, 0.7065735, 0.72592163, 0.7437744, 0.7574768, 0.7702942, 0.79159546, 0.82211304, 0.84536743, 0.84317017, 0.81503296, 0.783844, 0.7694092, 0.7713318, 0.7767029, 0.7810974, 0.7915039, 0.8138428, 0.84265137, 0.8684082, 0.89016724, 0.9100342, 0.9269409, 0.9324341, 0.92440796, 0.91519165, 0.91955566, 0.93881226, 0.9534607, 0.9517212, 0.9483032, 0.94589233, 0.9362793, 0.9331665, 0.94351196, 0.9481201, 0.9451904, 0.9439087, 0.94403076, 0.9426575, 0.9401245, 0.9293823, 0.91033936, 0.8980408, 0.90045166, 0.9158325, 0.9336548, 0.9397583, 0.93515015, 0.93460083, 0.93536377, 0.93325806, 0.9333801, 0.9326172, 0.93118286, 0.9313965, 0.93118286, 0.9313965,
		},
	}
	pluginPath string
)

const (
	zeroesProportionThreshold = 0
	sampleRate                = 44100.0
)

func init() {
	switch os := runtime.GOOS; os {
	case "darwin":
		// pluginPath = "_testdata/Krush.vst"
		pluginPath = "_testdata/TAL-Reverb.vst"
	case "windows":
		pluginPath = "_testdata/TAL-Reverb.dll"
	default:
		pluginPath = ""
	}
}

// sTest load plugin
func TestPlugin(t *testing.T) {
	library, err := Open(pluginPath)
	if err != nil {
		t.Fatalf("Failed to open library: %v\n", err)
	}
	defer library.Close()
	assert.NotNil(t, library.entryPoint)
	assert.NotNil(t, library.library)
	assert.NotNil(t, library.Name)
	assert.NotNil(t, library.Path)

	plugin, err := library.Open()
	if err != nil {
		t.Fatalf("Failed to open plugin: %v\n", err)
	}
	defer plugin.Close()
	assert.Equal(t, len(plugins), 1)
	assert.NotNil(t, plugin.effect)
	assert.NotNil(t, plugin.Name)
	assert.NotNil(t, plugin.Path)
	assert.Equal(t, true, plugin.CanProcessFloat32())

	plugin.Dispatch(EffOpen, 0, 0, nil, 0.0)

	// Set default sample rate and block size
	blocksize := int64(len(samples32[0]))
	plugin.Dispatch(EffSetSampleRate, 0, 0, nil, sampleRate)
	plugin.Dispatch(EffSetBlockSize, 0, blocksize, nil, 0.0)
	plugin.Dispatch(EffMainsChanged, 0, 1, nil, 0.0)

	if plugin.CanProcessFloat64() {
		assert.Equal(t, false, plugin.CanProcessFloat32())
		ps := plugin.ProcessFloat64(samples64)
		assert.NotNil(t, ps)
		assert.NotEmpty(t, ps)
		assert.Equal(t, len(samples64), len(ps))
		for c, s := range ps {
			assert.Equal(t, len(samples64[c]), len(s), "Output channel %v has wrong size. Expected: %v got: %v", c, len(samples64[c]), len(s))
			zc, zp, zpos := zeroesFloat64(s)
			assert.Equal(t, true, zeroesProportionThreshold >= zp, "Too many zeroed samples in channel %v expected: %v%% got: %.4f%% zeroes count: %v zeroes positions: %v", c, zeroesProportionThreshold, zp, zc, zpos)
		}
	}

	if plugin.CanProcessFloat32() {
		assert.Equal(t, false, plugin.CanProcessFloat64())
		ps := plugin.ProcessFloat32(samples32)
		assert.NotNil(t, ps)
		assert.NotEmpty(t, ps)
		assert.Equal(t, len(samples32), len(ps))
		for c, s := range ps {
			assert.Equal(t, len(samples32[c]), len(s), "Output channel %v has wrong size. Expected: %v got: %v", c, len(samples32[c]), len(s))
			zc, zp, zpos := zeroesFloat32(s)
			assert.Equal(t, true, zeroesProportionThreshold >= zp, "Too many zeroed samples in channel %v. Expected: %v%% got: %.4f%% zeroes count: %v zeroes positions: %v", c, zeroesProportionThreshold, zp, zc, zpos)
		}
	}
}

// count zeroes proportion in float64 slice
func zeroesFloat64(nums []float64) (count int, proportion float64, positions []int) {
	if nums == nil {
		return
	}

	positions = make([]int, 0, len(nums))
	for i, v := range nums {
		if v == 0.0 {
			count++
			positions = append(positions, i)
		}
	}
	proportion = float64(100*count) / float64(len(nums))
	return
}

// count zeroes proportion in float64 slice
func zeroesFloat32(nums []float32) (count int, proportion float64, positions []int) {
	if nums == nil {
		return
	}

	positions = make([]int, 0, len(nums))
	for i, v := range nums {
		if v == 0.0 {
			count++
			positions = append(positions, i)
		}
	}
	proportion = float64(100*count) / float64(len(nums))
	return
}
