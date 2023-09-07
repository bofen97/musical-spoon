package arxiv

/*
这里维护一个Topics [code subtopics 或许写到数据库]
*/
type SubTopic struct {
	Info         string
	Code         string
	SubTopicName string
}

type Topic struct {
	TopicName string
	SubTopics []*SubTopic
}

var Topics map[string]Topic

func init() {

	Topics = make(map[string]Topic)

	Topics["Computer Science"] = Topic{
		TopicName: "Computer Science",
		SubTopics: []*SubTopic{
			{
				Code:         "cs.ai",
				SubTopicName: "Artificial Intelligence",
				Info:         "Covers all areas of AI except Vision, Robotics, Machine Learning, Multiagent Systems, and Computation and Language (Natural Language Processing), which have separate subject areas. In particular, includes Expert Systems, Theorem Proving (although this may overlap with Logic in Computer Science), Knowledge Representation, Planning, and Uncertainty in AI. Roughly includes material in ACM Subject Classes I.2.0, I.2.1, I.2.3, I.2.4, I.2.8, and I.2.11.",
			},
			{
				Code:         "cs.ar",
				SubTopicName: "Hardware Architecture",
				Info:         "Covers systems organization and hardware architecture. Roughly includes material in ACM Subject Classes C.0, C.1, and C.5.",
			},
			{
				Code:         "cs.cc",
				SubTopicName: "Computational Complexity",
				Info:         "Covers models of computation, complexity classes, structural complexity, complexity tradeoffs, upper and lower bounds. Roughly includes material in ACM Subject Classes F.1 (computation by abstract devices), F.2.3 (tradeoffs among complexity measures), and F.4.3 (formal languages), although some material in formal languages may be more appropriate for Logic in Computer Science. Some material in F.2.1 and F.2.2, may also be appropriate here, but is more likely to have Data Structures and Algorithms as the primary subject area.",
			},
			{
				Code:         "cs.ce",
				SubTopicName: "Computational Engineering, Finance, and Science",
				Info:         "Covers applications of computer science to the mathematical modeling of complex systems in the fields of science, engineering, and finance. Papers here are interdisciplinary and applications-oriented, focusing on techniques and tools that enable challenging computational simulations to be performed, for which the use of supercomputers or distributed computing platforms is often required. Includes material in ACM Subject Classes J.2, J.3, and J.4 (economics).",
			},
			{
				Code:         "cs.cg",
				SubTopicName: "Computational Geometry",
				Info:         "Roughly includes material in ACM Subject Classes I.3.5 and F.2.2.",
			},
			{
				Code:         "cs.cl",
				SubTopicName: "Computation and Language",
				Info:         "Covers natural language processing. Roughly includes material in ACM Subject Class I.2.7. Note that work on artificial languages (programming languages, logics, formal systems) that does not explicitly address natural-language issues broadly construed (natural-language processing, computational linguistics, speech, text retrieval, etc.) is not appropriate for this area.",
			},
			{
				Code:         "cs.cr",
				SubTopicName: "Cryptography and Security",
				Info:         "Covers all areas of cryptography and security including authentication, public key cryptosytems, proof-carrying code, etc. Roughly includes material in ACM Subject Classes D.4.6 and E.3.",
			},
			{
				Code:         "cs.cv",
				SubTopicName: "Computer Vision and Pattern Recognition",
				Info:         "Covers image processing, computer vision, pattern recognition, and scene understanding. Roughly includes material in ACM Subject Classes I.2.10, I.4, and I.5.				",
			},

			{
				Code:         "cs.cy",
				SubTopicName: "Computers and Society",
				Info:         "Covers impact of computers on society, computer ethics, information technology and public policy, legal aspects of computing, computers and education. Roughly includes material in ACM Subject Classes K.0, K.2, K.3, K.4, K.5, and K.7.",
			},
			{
				Code:         "cs.db",
				SubTopicName: "Database",
				Info:         "Covers database management, datamining, and data processing. Roughly includes material in ACM Subject Classes E.2, E.5, H.0, H.2, and J.1.",
			},
			{
				Code:         "cs.dc",
				SubTopicName: "Distributed, Parallel, and Cluster Computing",
				Info:         "Covers fault-tolerance, distributed algorithms, stabilility, parallel computation, and cluster computing. Roughly includes material in ACM Subject Classes C.1.2, C.1.4, C.2.4, D.1.3, D.4.5, D.4.7, E.1.",
			},
			{
				Code:         "cs.dl",
				SubTopicName: "Digital Libraries",
				Info:         "Covers all aspects of the digital library design and document and text creation. Note that there will be some overlap with Information Retrieval (which is a separate subject area). Roughly includes material in ACM Subject Classes H.3.5, H.3.6, H.3.7, I.7.				",
			},
			{
				Code:         "cs.dm",
				SubTopicName: "Discrete Mathematics",
				Info:         "Covers combinatorics, graph theory, applications of probability. Roughly includes material in ACM Subject Classes G.2 and G.3.",
			},
			{
				Code:         "cs.ds",
				SubTopicName: "Structures and Algorithms",
				Info:         "Covers data structures and analysis of algorithms. Roughly includes material in ACM Subject Classes E.1, E.2, F.2.1, and F.2.2.				",
			},
			{
				Code:         "cs.et",
				SubTopicName: "Emerging Technologies",
				Info:         "Covers approaches to information processing (computing, communication, sensing) and bio-chemical analysis based on alternatives to silicon CMOS-based technologies, such as nanoscale electronic, photonic, spin-based, superconducting, mechanical, bio-chemical and quantum technologies (this list is not exclusive). Topics of interest include (1) building blocks for emerging technologies, their scalability and adoption in larger systems, including integration with traditional technologies, (2) modeling, design and optimization of novel devices and systems, (3) models of computation, algorithm design and programming for emerging technologies.",
			},
			{
				Code:         "cs.fl",
				SubTopicName: "Formal Languages and Automata Theory",
				Info:         "Covers automata theory, formal language theory, grammars, and combinatorics on words. This roughly corresponds to ACM Subject Classes F.1.1, and F.4.3. Papers dealing with computational complexity should go to cs.CC; papers dealing with logic should go to cs.LO.",
			},
			{
				Code:         "cs.gl",
				SubTopicName: "General Literature",
				Info:         "Covers introductory material, survey material, predictions of future trends, biographies, and miscellaneous computer-science related material. Roughly includes all of ACM Subject Class A, except it does not include conference proceedings (which will be listed in the appropriate subject area).",
			},
			{
				Code:         "cs.gr",
				SubTopicName: "Graphics",
				Info:         "Covers all aspects of computer graphics. Roughly includes material in all of ACM Subject Class I.3, except that I.3.5 is is likely to have Computational Geometry as the primary subject area.",
			},
			{
				Code:         "cs.gt",
				SubTopicName: "Computer Science and Game Theory",
				Info:         "Covers all theoretical and applied aspects at the intersection of computer science and game theory, including work in mechanism design, learning in games (which may overlap with Learning), foundations of agent modeling in games (which may overlap with Multiagent systems), coordination, specification and formal methods for non-cooperative computational environments. The area also deals with applications of game theory to areas such as electronic commerce.",
			},
			{
				Code:         "cs.hc",
				SubTopicName: "Human-Computer Interaction",
				Info:         "Covers human factors, user interfaces, and collaborative computing. Roughly includes material in ACM Subject Classes H.1.2 and all of H.5, except for H.5.1, which is more likely to have Multimedia as the primary subject area.",
			},
			{
				Code:         "cs.ir",
				SubTopicName: "Information Retrieval",
				Info:         "Covers indexing, dictionaries, retrieval, content and analysis. Roughly includes material in ACM Subject Classes H.3.0, H.3.1, H.3.2, H.3.3, and H.3.4.",
			},
			{
				Code:         "cs.it",
				SubTopicName: "Information Theory",
				Info:         "Covers theoretical and experimental aspects of information theory and coding. Includes material in ACM Subject Class E.4 and intersects with H.1.1.",
			},
			{
				Code:         "cs.lg",
				SubTopicName: "Machine Learning",
				Info:         "Papers on all aspects of machine learning research (supervised, unsupervised, reinforcement learning, bandit problems, and so on) including also robustness, explanation, fairness, and methodology. cs.LG is also an appropriate primary category for applications of machine learning methods.				",
			},
			{
				Code:         "cs.lo",
				SubTopicName: "Logic in Computer Science",
				Info:         "Covers all aspects of logic in computer science, including finite model theory, logics of programs, modal logic, and program verification. Programming language semantics should have Programming Languages as the primary subject area. Roughly includes material in ACM Subject Classes D.2.4, F.3.1, F.4.0, F.4.1, and F.4.2; some material in F.4.3 (formal languages) may also be appropriate here, although Computational Complexity is typically the more appropriate subject area.",
			},
			{
				Code:         "cs.ma",
				SubTopicName: "Multiagent Systems",
				Info:         "Covers multiagent systems, distributed artificial intelligence, intelligent agents, coordinated interactions. and practical applications. Roughly covers ACM Subject Class I.2.11.",
			},
			{
				Code:         "cs.mm",
				SubTopicName: "Multimedia",
				Info:         "Roughly includes material in ACM Subject Class H.5.1.",
			},
			{
				Code:         "cs.ms",
				SubTopicName: "Mathematical Software",
				Info:         "Roughly includes material in ACM Subject Class G.4.",
			},
			{
				Code:         "cs.na",
				SubTopicName: "Numerical Analysis",
				Info:         "cs.NA is an alias for math.NA. Roughly includes material in ACM Subject Class G.1.",
			},
			{
				Code:         "cs.ne",
				SubTopicName: "Neural and Evolutionary Computing",
				Info:         "Covers neural networks, connectionism, genetic algorithms, artificial life, adaptive behavior. Roughly includes some material in ACM Subject Class C.1.3, I.2.6, I.5.",
			},
			{
				Code:         "cs.ni",
				SubTopicName: "Networking and Internet Architecture",
				Info:         "Covers all aspects of computer communication networks, including network architecture and design, network protocols, and internetwork standards (like TCP/IP). Also includes topics, such as web caching, that are directly relevant to Internet architecture and performance. Roughly includes all of ACM Subject Class C.2 except C.2.4, which is more likely to have Distributed, Parallel, and Cluster Computing as the primary subject area.",
			},
			{
				Code:         "cs.oh",
				SubTopicName: "Other Computer Science",
				Info:         "This is the classification to use for documents that do not fit anywhere else.",
			},
			{
				Code:         "cs.os",
				SubTopicName: "Operating Systems",
				Info:         "Roughly includes material in ACM Subject Classes D.4.1, D.4.2., D.4.3, D.4.4, D.4.5, D.4.7, and D.4.9.",
			},
			{
				Code:         "cs.pf",
				SubTopicName: "Performance",
				Info:         "Covers performance measurement and evaluation, queueing, and simulation. Roughly includes material in ACM Subject Classes D.4.8 and K.6.2.",
			},
			{
				Code:         "cs.pl",
				SubTopicName: "Programming Languages",
				Info:         "Covers programming language semantics, language features, programming approaches (such as object-oriented programming, functional programming, logic programming). Also includes material on compilers oriented towards programming languages; other material on compilers may be more appropriate in Architecture (AR). Roughly includes material in ACM Subject Classes D.1 and D.3.",
			},
			{
				Code:         "cs.ro",
				SubTopicName: "robotics",
				Info:         "Roughly includes material in ACM Subject Class I.2.9.",
			},
			{
				Code:         "cs.sc",
				SubTopicName: "Symbolic Computation",
				Info:         "Roughly includes material in ACM Subject Class I.1.",
			},
			{
				Code:         "cs.sd",
				SubTopicName: "sound",
				Info:         "Covers all aspects of computing with sound, and sound as an information channel. Includes models of sound, analysis and synthesis, audio user interfaces, sonification of data, computer music, and sound signal processing. Includes ACM Subject Class H.5.5, and intersects with H.1.2, H.5.1, H.5.2, I.2.7, I.5.4, I.6.3, J.5, K.4.2.",
			},
			{
				Code:         "cs.se",
				SubTopicName: "Software Engineering",
				Info:         "Covers design tools, software metrics, testing and debugging, programming environments, etc. Roughly includes material in all of ACM Subject Classes D.2, except that D.2.4 (program verification) should probably have Logics in Computer Science as the primary subject area.",
			},
			{
				Code:         "cs.si",
				SubTopicName: "Social and Information Networks",
				Info:         "Covers the design, analysis, and modeling of social and information networks, including their applications for on-line information access, communication, and interaction, and their roles as datasets in the exploration of questions in these and other domains, including connections to the social and biological sciences. Analysis and modeling of such networks includes topics in ACM Subject classes F.2, G.2, G.3, H.2, and I.2; applications in computing include topics in H.3, H.4, and H.5; and applications at the interface of computing and other disciplines include topics in J.1--J.7. Papers on computer communication systems and network protocols (e.g. TCP/IP) are generally a closer fit to the Networking and Internet Architecture (cs.NI) category.",
			},
			{
				Code:         "cs.sy",
				SubTopicName: "Systems and Control",
				Info:         "cs.SY is an alias for eess.SY. This section includes theoretical and experimental research covering all facets of automatic control systems. The section is focused on methods of control system analysis and design using tools of modeling, simulation and optimization. Specific areas of research include nonlinear, distributed, adaptive, stochastic and robust control in addition to hybrid and discrete event systems. Application areas include automotive and aerospace control systems, network control, biological systems, multiagent and cooperative control, robotics, reinforcement learning, sensor networks, control of cyber-physical and energy-related systems, and control of computing systems.",
			},
		},
	}

	Topics["Economics"] = Topic{
		TopicName: "Economics",
		SubTopics: []*SubTopic{
			{
				Code:         "econ.em",
				SubTopicName: "Econometrics",
				Info:         "Econometric Theory, Micro-Econometrics, Macro-Econometrics, Empirical Content of Economic Relations discovered via New Methods, Methodological Aspects of the Application of Statistical Inference to Economic Data.",
			},
			{
				Code:         "econ.gn",
				SubTopicName: "General Economics",
				Info:         "General methodological, applied, and empirical contributions to economics.",
			},
			{
				Code:         "econ.th",
				SubTopicName: "Theoretical Economics",
				Info:         "Includes theoretical contributions to Contract Theory, Decision Theory, Game Theory, General Equilibrium, Growth, Learning and Evolution, Macroeconomics, Market and Mechanism Design, and Social Choice.",
			},
		},
	}

	Topics["Electrical Engineering and Systems Science"] = Topic{
		TopicName: "Electrical Engineering and Systems Science",
		SubTopics: []*SubTopic{
			{
				Code:         "eess.as",
				SubTopicName: "Audio and Speech Processing",
				Info:         "Theory and methods for processing signals representing audio, speech, and language, and their applications. This includes analysis, synthesis, enhancement, transformation, classification and interpretation of such signals as well as the design, development, and evaluation of associated signal processing systems. Machine learning and pattern analysis applied to any of the above areas is also welcome. Specific topics of interest include: auditory modeling and hearing aids; acoustic beamforming and source localization; classification of acoustic scenes; speaker separation; active noise control and echo cancellation; enhancement; de-reverberation; bioacoustics; music signals analysis, synthesis and modification; music information retrieval; audio for multimedia and joint audio-video processing; spoken and written language modeling, segmentation, tagging, parsing, understanding, and translation; text mining; speech production, perception, and psychoacoustics; speech analysis, synthesis, and perceptual modeling and coding; robust speech recognition; speaker recognition and characterization; deep learning, online learning, and graphical models applied to speech, audio, and language signals; and implementation aspects ranging from system architecture to fast algorithms.",
			},
			{
				Code:         "eess.iv",
				SubTopicName: "Image and Video Processing",
				Info:         "Theory, algorithms, and architectures for the formation, capture, processing, communication, analysis, and display of images, video, and multidimensional signals in a wide variety of applications. Topics of interest include: mathematical, statistical, and perceptual image and video modeling and representation; linear and nonlinear filtering, de-blurring, enhancement, restoration, and reconstruction from degraded, low-resolution or tomographic data; lossless and lossy compression and coding; segmentation, alignment, and recognition; image rendering, visualization, and printing; computational imaging, including ultrasound, tomographic and magnetic resonance imaging; and image and video analysis, synthesis, storage, search and retrieval.",
			},
			{
				Code:         "eess.SP",
				SubTopicName: "Signal Processing",
				Info:         "Theory, algorithms, performance analysis and applications of signal and data analysis, including physical modeling, processing, detection and parameter estimation, learning, mining, retrieval, and information extraction. The term \"signal\" includes speech, audio, sonar, radar, geophysical, physiological, (bio-) medical, image, video, and multimodal natural and man-made signals, including communication signals and data. Topics of interest include: statistical signal processing, spectral estimation and system identification; filter design, adaptive filtering / stochastic learning; (compressive) sampling, sensing, and transform-domain methods including fast algorithms; signal processing for machine learning and machine learning for signal processing applications; in-network and graph signal processing; convex and nonconvex optimization methods for signal processing applications; radar, sonar, and sensor array beamforming and direction finding; communications signal processing; low power, multi-core and system-on-chip signal processing; sensing, communication, analysis and optimization for cyber-physical systems such as power grids and the Internet of Things.",
			},
			{
				Code:         "eess.sy",
				SubTopicName: "Systems and Control",
				Info:         "This section includes theoretical and experimental research covering all facets of automatic control systems. The section is focused on methods of control system analysis and design using tools of modeling, simulation and optimization. Specific areas of research include nonlinear, distributed, adaptive, stochastic and robust control in addition to hybrid and discrete event systems. Application areas include automotive and aerospace control systems, network control, biological systems, multiagent and cooperative control, robotics, reinforcement learning, sensor networks, control of cyber-physical and energy-related systems, and control of computing systems.",
			},
		},
	}

	Topics["Mathematics"] = Topic{
		TopicName: "Mathematics",
		SubTopics: []*SubTopic{
			{
				Code:         "math.ac",
				SubTopicName: "Commutative Algebra",
				Info:         "Commutative rings, modules, ideals, homological algebra, computational aspects, invariant theory, connections to algebraic geometry and combinatorics",
			},
			{
				Code:         "math.ag",
				SubTopicName: "Algebraic Geometry",
				Info:         "Algebraic varieties, stacks, sheaves, schemes, moduli spaces, complex geometry, quantum cohomology",
			},
			{
				Code:         "math.ap",
				SubTopicName: "Analysis of PDEs",
				Info:         "Existence and uniqueness, boundary conditions, linear and non-linear operators, stability, soliton theory, integrable PDE's, conservation laws, qualitative dynamics",
			},
			{
				Code:         "math.at",
				SubTopicName: "Algebraic Topology",
				Info:         "Homotopy theory, homological algebra, algebraic treatments of manifolds",
			},
			{
				Code:         "math.ca",
				SubTopicName: "Classical Analysis and ODEs",
				Info:         "Special functions, orthogonal polynomials, harmonic analysis, ODE's, differential relations, calculus of variations, approximations, expansions, asymptotics",
			},
			{
				Code:         "math.co",
				SubTopicName: "Combinatorics",
				Info:         "Discrete mathematics, graph theory, enumeration, combinatorial optimization, Ramsey theory, combinatorial game theory",
			},
			{
				Code:         "math.ct",
				SubTopicName: "Category Theory",
				Info:         "Enriched categories, topoi, abelian categories, monoidal categories, homological algebra",
			},
			{
				Code:         "math.cv",
				SubTopicName: "Complex Variables",
				Info:         "Holomorphic functions, automorphic group actions and forms, pseudoconvexity, complex geometry, analytic spaces, analytic sheaves",
			},
			{
				Code:         "math.dg",
				SubTopicName: "Differential Geometry",
				Info:         "Complex, contact, Riemannian, pseudo-Riemannian and Finsler geometry, relativity, gauge theory, global analysis",
			},
			{
				Code:         "math.ds",
				SubTopicName: "Dynamical Systems",
				Info:         "Dynamics of differential equations and flows, mechanics, classical few-body problems, iterations, complex dynamics, delayed differential equations",
			},
			{
				Code:         "math.fa",
				SubTopicName: "Functional Analysis",
				Info:         "Banach spaces, function spaces, real functions, integral transforms, theory of distributions, measure theory",
			},
			{
				Code:         "math.gm",
				SubTopicName: "General Mathematics",
				Info:         "Mathematical material of general interest, topics not covered elsewhere",
			},
			{
				Code:         "math.gn",
				SubTopicName: "General Topology",
				Info:         "Continuum theory, point-set topology, spaces with algebraic structure, foundations, dimension theory, local and global properties",
			},
			{
				Code:         "math.gr",
				SubTopicName: "Group Theory",
				Info:         "Finite groups, topological groups, representation theory, cohomology, classification and structure",
			},
			{
				Code:         "math.gt",
				SubTopicName: "Geometric Topology",
				Info:         "Manifolds, orbifolds, polyhedra, cell complexes, foliations, geometric structures",
			},
			{
				Code:         "math.ho",
				SubTopicName: "History and Overview",
				Info:         "Biographies, philosophy of mathematics, mathematics education, recreational mathematics, communication of mathematics, ethics in mathematics",
			},
			{
				Code:         "math.it",
				SubTopicName: "Information Theory",
				Info:         "math.IT is an alias for cs.IT. Covers theoretical and experimental aspects of information theory and coding.",
			},
			{
				Code:         "math.kt",
				SubTopicName: "Theory and Homology",
				Info:         "Algebraic and topological K-theory, relations with topology, commutative algebra, and operator algebras",
			},
			{
				Code:         "math.lo",
				SubTopicName: "Logic",
				Info:         "Logic, set theory, point-set topology, formal mathematics",
			},
			{
				Code:         "math.mg",
				SubTopicName: "Metric Geometry",
				Info:         "Euclidean, hyperbolic, discrete, convex, coarse geometry, comparisons in Riemannian geometry, symmetric spaces",
			},
			{
				Code:         "math.mp",
				SubTopicName: "Mathematical Physics",
				Info:         "math.MP is an alias for math-ph. Articles in this category focus on areas of research that illustrate the application of mathematics to problems in physics, develop mathematical methods for such applications, or provide mathematically rigorous formulations of existing physical theories. Submissions to math-ph should be of interest to both physically oriented mathematicians and mathematically oriented physicists; submissions which are primarily of interest to theoretical physicists or to mathematicians should probably be directed to the respective physics/math categories",
			},
			{
				Code:         "math.na",
				SubTopicName: "Numerical Analysis",
				Info:         "Numerical algorithms for problems in analysis and algebra, scientific computation",
			},
			{
				Code:         "math.nt",
				SubTopicName: "Number Theory",
				Info:         "Prime numbers, diophantine equations, analytic number theory, algebraic number theory, arithmetic geometry, Galois theory",
			},
			{
				Code:         "math.oa",
				SubTopicName: "Operator Algebras",
				Info:         "Algebras of operators on Hilbert space, C^*-algebras, von Neumann algebras, non-commutative geometry",
			},
			{
				Code:         "math.oc",
				SubTopicName: "Optimization and Control",
				Info:         "Operations research, linear programming, control theory, systems theory, optimal control, game theory",
			},
			{
				Code:         "math.pr",
				SubTopicName: "Probability",
				Info:         "Theory and applications of probability and stochastic processes: e.g. central limit theorems, large deviations, stochastic differential equations, models from statistical mechanics, queuing theory",
			},
			{
				Code:         "math.qa",
				SubTopicName: "Quantum Algebra",
				Info:         "Quantum groups, skein theories, operadic and diagrammatic algebra, quantum field theory",
			},
			{
				Code:         "math.ra",
				SubTopicName: "Rings and Algebras",
				Info:         "Non-commutative rings and algebras, non-associative algebras, universal algebra and lattice theory, linear algebra, semigroups",
			},
			{
				Code:         "math.rt",
				SubTopicName: "Representation Theory",
				Info:         "Linear representations of algebras and groups, Lie theory, associative algebras, multilinear algebra",
			},
			{
				Code:         "math.sg",
				SubTopicName: "Symplectic Geometry",
				Info:         "Hamiltonian systems, symplectic flows, classical integrable systems",
			},
			{
				Code:         "math.sp",
				SubTopicName: "Spectral Theory",
				Info:         "Schrodinger operators, operators on manifolds, general differential operators, numerical studies, integral operators, discrete models, resonances, non-self-adjoint operators, random operators/matrices",
			},
			{
				Code:         "math.st",
				SubTopicName: "Statistics Theory",
				Info:         "Applied, computational and theoretical statistics: e.g. statistical inference, regression, time series, multivariate analysis, data analysis, Markov chain Monte Carlo, design of experiments, case studies",
			},
		},
	}

	Topics["Physics"] = Topic{
		TopicName: "Physics",
		SubTopics: []*SubTopic{
			{
				Code:         "astro-ph.co",
				SubTopicName: "Cosmology and Nongalactic Astrophysics",
				Info:         "Phenomenology of early universe, cosmic microwave background, cosmological parameters, primordial element abundances, extragalactic distance scale, large-scale structure of the universe. Groups, superclusters, voids, intergalactic medium. Particle astrophysics: dark energy, dark matter, baryogenesis, leptogenesis, inflationary models, reheating, monopoles, WIMPs, cosmic strings, primordial black holes, cosmological gravitational radiation",
			},
			{
				Code:         "astro-ph.ga",
				SubTopicName: "Astrophysics of Galaxies",
				Info:         "Phenomena pertaining to galaxies or the Milky Way. Star clusters, HII regions and planetary nebulae, the interstellar medium, atomic and molecular clouds, dust. Stellar populations. Galactic structure, formation, dynamics. Galactic nuclei, bulges, disks, halo. Active Galactic Nuclei, supermassive black holes, quasars. Gravitational lens systems. The Milky Way and its contents",
			},
			{
				Code:         "astro-ph.ep",
				SubTopicName: "Earth and Planetary Astrophysics",
				Info:         "Interplanetary medium, planetary physics, planetary astrobiology, extrasolar planets, comets, asteroids, meteorites. Structure and formation of the solar system",
			},
			{
				Code:         "astro-ph.he",
				SubTopicName: "High Energy Astrophysical Phenomena",
				Info:         "Cosmic ray production, acceleration, propagation, detection. Gamma ray astronomy and bursts, X-rays, charged particles, supernovae and other explosive phenomena, stellar remnants and accretion systems, jets, microquasars, neutron stars, pulsars, black holes",
			},
			{
				Code:         "astro-ph.im",
				SubTopicName: "Instrumentation and Methods for Astrophysics",
				Info:         "Detector and telescope design, experiment proposals. Laboratory Astrophysics. Methods for data analysis, statistical methods. Software, database design",
			},
			{
				Code:         "astro-ph.sr",
				SubTopicName: "Solar and Stellar Astrophysics",
				Info:         "White dwarfs, brown dwarfs, cataclysmic variables. Star formation and protostellar systems, stellar astrobiology, binary and multiple systems of stars, stellar evolution and structure, coronas. Central stars of planetary nebulae. Helioseismology, solar neutrinos, production and detection of gravitational radiation from stellar systems",
			},
			{
				Code:         "cond-mat.dis-nn",
				SubTopicName: "Disordered Systems and Neural Networks",
				Info:         "Glasses and spin glasses; properties of random, aperiodic and quasiperiodic systems; transport in disordered media; localization; phenomena mediated by defects and disorder; neural networks",
			},
			{
				Code:         "cond-mat.mes-hall",
				SubTopicName: "Mesoscale and Nanoscale Physics",
				Info:         "Semiconducting nanostructures: quantum dots, wires, and wells. Single electronics, spintronics, 2d electron gases, quantum Hall effect, nanotubes, graphene, plasmonic nanostructures",
			},
			{
				Code:         "cond-mat.mtrl-sci",
				SubTopicName: "Materials Science",
				Info:         "Techniques, synthesis, characterization, structure. Structural phase transitions, mechanical properties, phonons. Defects, adsorbates, interfaces",
			},
			{

				Code:         "cond-mat.other",
				SubTopicName: "Other Condensed Matter",
				Info:         "Work in condensed matter that does not fit into the other cond-mat classifications",
			},
			{
				Code:         "cond-mat.quant-gas",
				SubTopicName: "Quantum Gases",
				Info:         "Ultracold atomic and molecular gases, Bose-Einstein condensation, Feshbach resonances, spinor condensates, optical lattices, quantum simulation with cold atoms and molecules, macroscopic interference phenomena",
			},
			{
				Code:         "cond-mat.soft",
				SubTopicName: "Soft Condensed Matter",
				Info:         "Membranes, polymers, liquid crystals, glasses, colloids, granular matter",
			},
			{
				Code:         "cond-mat.stat-mech",
				SubTopicName: "Statistical Mechanics",
				Info:         "Phase transitions, thermodynamics, field theory, non-equilibrium phenomena, renormalization group and scaling, integrable models, turbulence",
			},
			{
				Code:         "cond-mat.str-el",
				SubTopicName: "Strongly Correlated Electrons",
				Info:         "Quantum magnetism, non-Fermi liquids, spin liquids, quantum criticality, charge density waves, metal-insulator transitions",
			},
			{
				Code:         "cond-mat.supr-con",
				SubTopicName: "Superconductivity",

				Info: "Superconductivity: theory, models, experiment. Superflow in helium",
			},
			{
				Code:         "gr-qc",
				SubTopicName: "General Relativity and Quantum Cosmology",
				Info:         "General Relativity and Quantum Cosmology Areas of gravitational physics, including experiments and observations related to the detection and interpretation of gravitational waves, experimental tests of gravitational theories, computational general relativity, relativistic astrophysics, solutions to Einstein's equations and their properties, alternative theories of gravity, classical and quantum cosmology, and quantum gravity.",
			},
			{
				Code:         "hep-ex",
				SubTopicName: "High Energy Physics - Experiment",
				Info:         "Description coming soon",
			},
			{
				Code:         "hep-lat",
				SubTopicName: "High Energy Physics - Lattice",
				Info:         "Lattice field theory. Phenomenology from lattice field theory. Algorithms for lattice field theory. Hardware for lattice field theory.",
			},
			{
				Code:         "hep-ph",
				SubTopicName: "High Energy Physics - Phenomenology",
				Info:         "Theoretical particle physics and its interrelation with experiment. Prediction of particle physics observables: models, effective field theories, calculation techniques. Particle physics: analysis of theory through experimental results",
			},
			{
				Code:         "hep-th",
				SubTopicName: "High Energy Physics - Theory",
				Info:         "Formal aspects of quantum field theory. String theory, supersymmetry and supergravity.",
			},
			{
				Code:         "math-ph",
				SubTopicName: "Mathematical Physics",
				Info:         "Articles in this category focus on areas of research that illustrate the application of mathematics to problems in physics, develop mathematical methods for such applications, or provide mathematically rigorous formulations of existing physical theories. Submissions to math-ph should be of interest to both physically oriented mathematicians and mathematically oriented physicists; submissions which are primarily of interest to theoretical physicists or to mathematicians should probably be directed to the respective physics/math categories",
			},
			{
				Code:         "nlin.ao",
				SubTopicName: "Adaptation and Self-Organizing Systems",
				Info:         "Adaptation, self-organizing systems, statistical physics, fluctuating systems, stochastic processes, interacting particle systems, machine learning",
			},
			{
				Code:         "nlin.cd",
				SubTopicName: "Chaotic Dynamics",
				Info:         "Dynamical systems, chaos, quantum chaos, topological dynamics, cycle expansions, turbulence, propagation",
			},
			{
				Code:         "nlin.cg",
				SubTopicName: "Cellular Automata and Lattice Gases",
				Info:         "Computational methods, time series analysis, signal processing, wavelets, lattice gases",
			},
			{
				Code:         "nlin.ps",
				SubTopicName: "Pattern Formation and Solitons",
				Info:         "Pattern formation, coherent structures, solitons",
			},
			{
				Code:         "nlin.si",
				SubTopicName: "Exactly Solvable and Integrable Systems",
				Info:         "Exactly solvable systems, integrable PDEs, integrable ODEs, Painleve analysis, integrable discrete maps, solvable lattice models, integrable quantum systems",
			},
			{
				Code:         "nucl-ex",
				SubTopicName: "Nuclear Experiment",

				Info: "Nuclear Experiment Results from experimental nuclear physics including the areas of fundamental interactions, measurements at low- and medium-energy, as well as relativistic heavy-ion collisions. Does not include: detectors and instrumentation nor analysis methods to conduct experiments; descriptions of experimental programs (present or future); comments on published results",
			},
			{
				Code:         "nucl-th",
				SubTopicName: "Nuclear Theory",
				Info:         "Nuclear Theory Theory of nuclear structure covering wide area from models of hadron structure to neutron stars. Nuclear equation of states at different external conditions. Theory of nuclear reactions including heavy-ion reactions at low and high energies. It does not include problems of data analysis, physics of nuclear reactors, problems of safety, reactor construction",
			},
			{
				Code:         "physics.acc-ph",
				SubTopicName: "Accelerator Physics",
				Info:         "Accelerator theory and simulation. Accelerator technology. Accelerator experiments. Beam Physics. Accelerator design and optimization. Advanced accelerator concepts. Radiation sources including synchrotron light sources and free electron lasers. Applications of accelerators.",
			},
			{
				Code:         "physics.app-ph",
				SubTopicName: "Applied Physics",
				Info:         "Applications of physics to new technology, including electronic devices, optics, photonics, microwaves, spintronics, advanced materials, metamaterials, nanotechnology, and energy sciences.",
			},
			{
				Code:         "physics.atm-clus",
				SubTopicName: "Atomic and Molecular Clusters",
				Info:         "Atomic and molecular clusters, nanoparticles: geometric, electronic, optical, chemical, magnetic properties, shell structure, phase transitions, optical spectroscopy, mass spectrometry, photoelectron spectroscopy, ionization potential, electron affinity, interaction with intense light pulses, electron diffraction, light scattering, ab initio calculations, DFT theory, fragmentation, Coulomb explosion, hydrodynamic expansion.",
			},
			{
				Code:         "physics.atom-ph",
				SubTopicName: "Atomic Physics",
				Info:         "Atomic and molecular structure, spectra, collisions, and data. Atoms and molecules in external fields. Molecular dynamics and coherent and optical control. Cold atoms and molecules. Cold collisions. Optical lattices.",
			},
			{
				Code:         "physics.bio-ph",
				SubTopicName: "Biological Physics",
				Info:         "Molecular biophysics, cellular biophysics, neurological biophysics, membrane biophysics, single-molecule biophysics, ecological biophysics, quantum phenomena in biological systems (quantum biophysics), theoretical biophysics, molecular dynamics/modeling and simulation, game theory, biomechanics, bioinformatics, microorganisms, virology, evolution, biophysical methods.",
			},
			{
				Code:         "physics.chem-ph",
				SubTopicName: "Chemical Physics",
				Info:         "Experimental, computational, and theoretical physics of atoms, molecules, and clusters - Classical and quantum description of states, processes, and dynamics; spectroscopy, electronic structure, conformations, reactions, interactions, and phases. Chemical thermodynamics. Disperse systems. High pressure chemistry. Solid state chemistry. Surface and interface chemistry.",
			},
			{
				Code:         "physics.class-ph",
				SubTopicName: "Classical Physics",
				Info:         "Newtonian and relativistic dynamics; many particle systems; planetary motions; chaos in classical dynamics. Maxwell's equations and dynamics of charged systems and electromagnetic forces in materials. Vibrating systems such as membranes and cantilevers; optomechanics. Classical waves, including acoustics and elasticity; physics of music and musical instruments. Classical thermodynamics and heat flow problems.",
			},
			{
				Code:         "physics.comp-ph",
				SubTopicName: "Computational Physics",
				Info:         "All aspects of computational science applied to physics.",
			},
			{
				Code:         "physics.data-an",
				SubTopicName: "Data Analysis, Statistics and Probability",
				Info:         "Methods, software and hardware for physics data analysis: data processing and storage; measurement methodology; statistical and mathematical aspects such as parametrization and uncertainties.",
			},
			{
				Code:         "physics.ed-ph",
				SubTopicName: "Physics Education",
				Info:         "Report of results of a research study, laboratory experience, assessment or classroom practice that represents a way to improve teaching and learning in physics. Also, report on misconceptions of students, textbook errors, and other similar information relative to promoting physics understanding.",
			},
			{
				Code:         "physics.flu-dyn",
				SubTopicName: "Fluid Dynamics",
				Info:         "Turbulence, instabilities, incompressible/compressible flows, reacting flows. Aero/hydrodynamics, fluid-structure interactions, acoustics. Biological fluid dynamics, micro/nanofluidics, interfacial phenomena. Complex fluids, suspensions and granular flows, porous media flows. Geophysical flows, thermoconvective and stratified flows. Mathematical and computational methods for fluid dynamics, fluid flow models, experimental techniques.",
			},
			{
				Code:         "physics.gen-ph",
				SubTopicName: "General Physics",
				Info:         "Description coming soon",
			},
			{
				Code:         "physics.geo-ph",
				SubTopicName: "Geophysics",
				Info:         "Atmospheric physics. Biogeosciences. Computational geophysics. Geographic location. Geoinformatics. Geophysical techniques. Hydrospheric geophysics. Magnetospheric physics. Mathematical geophysics. Planetology. Solar system. Solid earth geophysics. Space plasma physics. Mineral physics. High pressure physics.",
			},
			{
				Code:         "physics.hist-ph",
				SubTopicName: "History and Philosophy of Physics",
				Info:         "History and philosophy of all branches of physics, astrophysics, and cosmology, including appreciations of physicists.",
			},
			{
				Code: "physics.ins-det",

				SubTopicName: "Instrumentation and Detectors",
				Info:         "Instrumentation and Detectors for research in natural science, including optical, molecular, atomic, nuclear and particle physics instrumentation and the associated electronics, services, infrastructure and control equipment.",
			},
			{
				Code:         "physics.med-ph",
				SubTopicName: "Medical Physics",
				Info:         "Radiation therapy. Radiation dosimetry. Biomedical imaging modelling. Reconstruction, processing, and analysis. Biomedical system modelling and analysis. Health physics. New imaging or therapy modalities.",
			},
			{
				Code:         "physics.optics",
				SubTopicName: "Optics",
				Info:         "Adaptive optics. Astronomical optics. Atmospheric optics. Biomedical optics. Cardinal points. Collimation. Doppler effect. Fiber optics. Fourier optics. Geometrical optics (Gradient index optics. Holography. Infrared optics. Integrated optics. Laser applications. Laser optical systems. Lasers. Light amplification. Light diffraction. Luminescence. Microoptics. Nano optics. Ocean optics. Optical computing. Optical devices. Optical imaging. Optical materials. Optical metrology. Optical microscopy. Optical properties. Optical signal processing. Optical testing techniques. Optical wave propagation. Paraxial optics. Photoabsorption. Photoexcitations. Physical optics. Physiological optics. Quantum optics. Segmented optics. Spectra. Statistical optics. Surface optics. Ultrafast optics. Wave optics. X-ray optics.",
			},
			{
				Code:         "physics.plasm-ph",
				SubTopicName: "Plasma Physics",
				Info:         "Fundamental plasma physics. Magnetically Confined Plasmas (includes magnetic fusion energy research). High Energy Density Plasmas (inertial confinement plasmas, laser-plasma interactions). Ionospheric, Heliophysical, and Astrophysical plasmas (includes sun and solar system plasmas). Lasers, Accelerators, and Radiation Generation. Low temperature plasmas and plasma applications (include dusty plasmas, semiconductor etching, plasma-based nanotechnology, medical applications). Plasma Diagnostics, Engineering and Enabling Technologies (includes fusion reactor design, heating systems, diagnostics, experimental techniques)",
			},
			{
				Code:         "physics.pop-ph",
				SubTopicName: "Popular Physics",
				Info:         "Description coming soon",
			},
			{
				Code:         "physics.soc-ph",
				SubTopicName: "Physics and Society",
				Info:         "Structure, dynamics and collective behavior of societies and groups (human or otherwise). Quantitative analysis of social networks and other complex networks. Physics and engineering of infrastructure and systems of broad societal impact (e.g., energy grids, transportation networks).",
			},
			{
				Code:         "physics.space-ph",
				SubTopicName: "Space Physics",
				Info:         "Space plasma physics. Heliophysics. Space weather. Planetary magnetospheres, ionospheres and magnetotail. Auroras. Interplanetary space. Cosmic rays. Synchrotron radiation. Radio astronomy.",
			},
			{
				Code:         "quant-ph",
				SubTopicName: "Quantum Physics",
				Info:         "Description coming soon",
			},
		},
	}

	Topics["Quantitative Biology"] = Topic{
		TopicName: "Quantitative Biology",
		SubTopics: []*SubTopic{
			{
				Code:         "q-bio.BM",
				SubTopicName: "Biomolecules",
				Info:         "DNA, RNA, proteins, lipids, etc.; molecular structures and folding kinetics; molecular interactions; single-molecule manipulation.",
			},
			{
				Code:         "q-bio.CB",
				SubTopicName: "Cell Behavior",
				Info:         "Cell-cell signaling and interaction; morphogenesis and development; apoptosis; bacterial conjugation; viral-host interaction; immunology",
			},
			{
				Code:         "q-bio.GN",
				SubTopicName: "Genomics",
				Info:         "DNA sequencing and assembly; gene and motif finding; RNA editing and alternative splicing; genomic structure and processes (replication, transcription, methylation, etc); mutational processes.",
			},
			{
				Code:         "q-bio.MN",
				SubTopicName: "Molecular Networks",
				Info:         "Gene regulation, signal transduction, proteomics, metabolomics, gene and enzymatic networks",
			},
			{
				Code:         "q-bio.NC",
				SubTopicName: "Neurons and Cognition",
				Info:         "Synapse, cortex, neuronal dynamics, neural network, sensorimotor control, behavior, attention",
			},
			{
				Code:         "q-bio.OT",
				SubTopicName: "Other Quantitative Biology",
				Info:         "Work in quantitative biology that does not fit into the other q-bio classifications",
			},
			{
				Code:         "q-bio.PE",
				SubTopicName: "Populations and Evolution",
				Info:         "Population dynamics, spatio-temporal and epidemiological models, dynamic speciation, co-evolution, biodiversity, foodwebs, aging; molecular evolution and phylogeny; directed evolution; origin of life",
			},
			{
				Code:         "q-bio.QM",
				SubTopicName: "Quantitative Methods",
				Info:         "All experimental, numerical, statistical and mathematical contributions of value to biology",
			},
			{
				Code:         "q-bio.SC",
				SubTopicName: "Subcellular Processes",
				Info:         "Assembly and control of subcellular structures (channels, organelles, cytoskeletons, capsules, etc.); molecular motors, transport, subcellular localization; mitosis and meiosis",
			},
			{
				Code:         "q-bio.TO",
				SubTopicName: "Tissues and Organs",
				Info:         "Blood flow in vessels, biomechanics of bones, electrical waves, endocrine system, tumor growth",
			},
		},
	}
	Topics["Quantitative Finance"] = Topic{
		TopicName: "Quantitative Finance",
		SubTopics: []*SubTopic{
			{
				Code:         "q-fin.CP",
				SubTopicName: "Computational Finance",
				Info:         "Computational methods, including Monte Carlo, PDE, lattice and other numerical methods with applications to financial modeling",
			},
			{
				Code:         "q-fin.EC",
				SubTopicName: "Economics",
				Info:         "q-fin.EC is an alias for econ.GN. Economics, including micro and macro economics, international economics, theory of the firm, labor economics, and other economic topics outside finance",
			},
			{
				Code:         "q-fin.GN",
				SubTopicName: "General Finance",
				Info:         "Development of general quantitative methodologies with applications in finance",
			},
			{
				Code:         "q-fin.MF",
				SubTopicName: "Mathematical Finance",
				Info:         "Mathematical and analytical methods of finance, including stochastic, probabilistic and functional analysis, algebraic, geometric and other methods",
			},
			{
				Code:         "q-fin.PM",
				SubTopicName: "Portfolio Management",
				Info:         "Security selection and optimization, capital allocation, investment strategies and performance measurement",
			},
			{
				Code:         "q-fin.PR",
				SubTopicName: "Pricing of Securities",
				Info:         "Valuation and hedging of financial securities, their derivatives, and structured products",
			},
			{
				Code:         "q-fin.RM",
				SubTopicName: "Risk Management",
				Info:         "Measurement and management of financial risks in trading, banking, insurance, corporate and other applications",
			},
			{
				Code:         "q-fin.ST",
				SubTopicName: "Statistical Finance",
				Info:         "Statistical, econometric and econophysics analyses with applications to financial markets and economic data",
			},

			{
				Code:         "q-fin.TR",
				SubTopicName: "Trading and Market Microstructure",
				Info:         "Market microstructure, liquidity, exchange and auction design, automated trading, agent-based modeling and market-making",
			},
		},
	}
	Topics["Statistics"] = Topic{
		TopicName: "Statistics",
		SubTopics: []*SubTopic{
			{
				Code:         "stat.AP",
				SubTopicName: "Applications",
				Info:         "Biology, Education, Epidemiology, Engineering, Environmental Sciences, Medical, Physical Sciences, Quality Control, Social Sciences",
			},
			{
				Code:         "stat.CO",
				SubTopicName: "Computation",
				Info:         "Algorithms, Simulation, Visualization",
			},
			{
				Code:         "stat.ME",
				SubTopicName: "Methodology",
				Info:         "Design, Surveys, Model Selection, Multiple Testing, Multivariate Methods, Signal and Image Processing, Time Series, Smoothing, Spatial Statistics, Survival Analysis, Nonparametric and Semiparametric Methods",
			},
			{
				Code:         "stat.ML",
				SubTopicName: "Machine Learning",
				Info:         "Covers machine learning papers (supervised, unsupervised, semi-supervised learning, graphical models, reinforcement learning, bandits, high dimensional inference, etc.) with a statistical or theoretical grounding",
			},
			{
				Code:         "stat.OT",
				SubTopicName: "Other Statistics",
				Info:         "Work in statistics that does not fit into the other stat classifications",
			},
			{
				Code:         "stat.TH",
				SubTopicName: "Statistics Theory",
				Info:         "stat.TH is an alias for math.ST. Asymptotics, Bayesian Inference, Decision Theory, Estimation, Foundations, Inference, Testing.",
			},
		},
	}
}
