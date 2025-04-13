package exsrapi

/*
	00AA00	FuncNameOfThisFunction()の取得対象関数レベルを変更し、テスト関数を作る（バージョン付与開始）
	00AB00	獲得ポイント取得対象ルームの登録で足切りの基準を定める（EventInfのThinitとThdelta）　SetThdata()、GetEventinf()
			MakeSampleTime() データ取得タイミング（分、秒）を生成する
			FuncNameOfThisFunction() 引数で自分自身の名称、親の名称のいずれかを取得できるようにする
			PrtHdr() FuncNameOfThisFunction() のインターフェース変更に対応する
	00AB01	PrtHdr()とPrintExf()の書式を統一する
	00AC01	GetEventidOfEventBox()であたらしいボックスイベントページに対応する。GetEventinf()でAPIを利用して情報を取得する
	00AD00	CheckExistingLock()を追加する。ロックファイルの存在確認とプロセスの存在確認を行う
*/

const Version="00AD00"
